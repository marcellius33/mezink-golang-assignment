package routers

import (
	"mezink-golang-assignment/controllers"

	"github.com/gin-gonic/gin"
)

func InitRecordRoutes(Routes *gin.Engine, controller *controllers.RecordController) {
	recordRouter := Routes.Group("/records")
	{
		recordRouter.POST("", controller.GetRecords)
	}
}
