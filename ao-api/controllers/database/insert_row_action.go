package database

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Usage:

	curl -X POST -H "Content-Type: application/json" -d '{
		"column_name": "column_value"
	}' http://localhost:3004/database/query/insert/project/<project_tab>/table/<table_name>


*/

func (dc *DatabaseController) InsertRow() gin.HandlerFunc {
	return func(c *gin.Context) {

		projectTag := c.Param("project_tag")
		tableName := c.Param("table_name")

		var dto map[string]string
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		fmt.Println(dto)

		if err := dc.Service.InsertRow(projectTag, tableName, dto); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(200, gin.H{"message": "Row added successfully"})

	}
}
