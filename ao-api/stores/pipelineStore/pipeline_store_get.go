package pipelineStore

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"

	"github.com/dotenx/dotenx/ao-api/db"
	"github.com/dotenx/dotenx/ao-api/models"
)

func (p *pipelineStore) GetByName(context context.Context, accountId string, name string) (pipeline models.PipelineVersion, endpoint string, isActive bool, isTemplate bool, isInteraction bool, err error) {
	// In the future we can use different statements based on the db.Driver as per DB Engine
	pipeline.Manifest.Tasks = make(map[string]models.Task)

	switch p.db.Driver {
	case db.Postgres:
		conn := p.db.Connection
		err = conn.QueryRow(select_pipeline, accountId, name).Scan(&pipeline.Id, &endpoint, &isActive, &isTemplate, &isInteraction)
		if err != nil {
			if err == sql.ErrNoRows {
				err = errors.New("not found")
				return
			}
			log.Println("error", err.Error())
			return
		}
		tasks := []models.Task{}
		var rows *sql.Rows
		rows, err = conn.Query(select_tasks_by_pipeline_id, pipeline.Id)
		if err != nil {
			log.Println("error", err.Error())
			return
		}
		defer rows.Close()
		for rows.Next() {
			task := models.Task{}
			var body interface{}
			err = rows.Scan(&task.Id, &task.Name, &task.Type, &task.Integration, &task.Description, &body)
			if err != nil {
				return
			}
			var taskBody models.TaskBodyMap
			json.Unmarshal(body.([]byte), &taskBody)
			task.Body = taskBody
			task.MetaData = models.AvaliableTasks[task.Type]
			tasks = append(tasks, task)
		}
		taskIdToName := make(map[int]string)
		for _, task := range tasks {
			taskIdToName[task.Id] = task.Name
		}
		for _, task := range tasks {
			preconditions := []struct {
				PreconditionId int    `db:"precondition_id"`
				Status         string `db:"status"`
			}{}
			err = conn.Select(&preconditions, select_preconditions_by_task_id, task.Id)
			if err != nil {
				return
			}
			task.ExecuteAfter = make(map[string][]string)
			for _, precondition := range preconditions {
				task.ExecuteAfter[taskIdToName[precondition.PreconditionId]] = append(task.ExecuteAfter[taskIdToName[precondition.PreconditionId]], precondition.Status)
			}
			pipeline.Manifest.Tasks[task.Name] = models.Task{
				ExecuteAfter: task.ExecuteAfter,
				Type:         task.Type,
				Body:         task.Body,
				Description:  task.Description,
				Integration:  task.Integration,
				MetaData:     task.MetaData,
			}
		}
	}
	return pipeline, endpoint, isActive, isTemplate, isInteraction, nil
}

var select_pipeline = `
SELECT id , endpoint, is_active, is_template, is_interaction
FROM pipelines p
WHERE account_id = $1 AND name = $2
`
var select_tasks_by_pipeline_id = `
SELECT id, name, task_type, integration, description, body FROM tasks
WHERE pipeline_id = $1
`
var select_preconditions_by_task_id = `
SELECT precondition_id, status FROM task_preconditions
WHERE task_id = $1
`
