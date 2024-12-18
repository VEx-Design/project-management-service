package handler

import (
	"net/http"
	"project-management-service/external/handler/request"
	"project-management-service/internal/core/entities"
	"project-management-service/internal/core/logic"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	projSrv logic.ProjectService
}

func NewProjectHandler(projSrv logic.ProjectService) *ProjectHandler {
	return &ProjectHandler{projSrv: projSrv}
}

func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var request request.Project
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newProj := entities.Project{
		Name:        request.Name,
		Description: request.Description,
	}
	h.projSrv.CreateProject(newProj)
	c.JSON(200, nil)
}
