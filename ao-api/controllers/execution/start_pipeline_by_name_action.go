package execution

import (
	"log"
	"net/http"

	"github.com/dotenx/dotenx/ao-api/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (e *ExecutionController) StartPipelineByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		accountId, _ := utils.GetAccountId(c)

		name := c.Param("name")
		// Get the `input data` from the request body
		var input map[string]interface{}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		res, err := e.Service.StartPipelineByName(input, accountId, name)
		if err != nil {
			if err.Error() == "automation is not active" {
				c.JSON(http.StatusBadRequest, err.Error())
				return
			} else if err.Error() == "no value for this field in initial data or return values" {
				c.JSON(http.StatusBadRequest, "some of your task fields are dependent on your trigger")
				return
			}
			log.Println(err.Error())
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, res)
	}
}
