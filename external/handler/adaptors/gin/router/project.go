package router

import (
	handler "project-management-service/external/handler/adaptors/gin/api"

	"github.com/gin-gonic/gin"
)

func RegisterProjectRoutes(router *gin.Engine, projHandler *handler.ProjectHandler) {
	api := router.Group("/api/v1")
	{
		api.POST("/project", projHandler.CreateProject)
		api.GET("/project", projHandler.GetMyProject)
	}
}
