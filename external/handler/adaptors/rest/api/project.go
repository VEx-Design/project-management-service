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

// CreateProject handles the creation of a new project.
func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var req request.Project
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	// Validate required fields
	if req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project name is required"})
		return
	}

	userID := c.GetHeader("X-User-Id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in headers"})
		return
	}

	// Create a new project entity
	newProj := entities.Project{
		Name:            req.Name,
		Description:     req.Description,
		OwnerId:         userID,
		Flow:            req.Flow,
		TypesConfig:     req.TypesConfig,
		ConfigurationID: req.ConfigurationID,
	}

	if err := h.projSrv.CreateProject(newProj); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Project created successfully"})
}

// GetMyProject retrieves the projects for the current user.
func (h *ProjectHandler) GetMyProject(c *gin.Context) {
	userId := c.Query("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	projects, err := h.projSrv.GetMyProject(c, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve projects", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"projects": projects})
}

// Update project name and description
func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	var req request.UpdateProject
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	// Validate required fields
	if req.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
		return
	}

	userID := c.GetHeader("X-User-Id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in headers"})
		return
	}

	updateProj := entities.Project{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
	}

	if err := h.projSrv.UpdateProject(updateProj, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project updated successfully"})
}

// UpdateProjectFlow updates the project flow
func (h *ProjectHandler) UpdateProjectFlow(c *gin.Context) {
	var req request.UpdateProjectFlow
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	// Validate required fields
	if req.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
		return
	}

	userID := c.GetHeader("X-User-Id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in headers"})
		return
	}

	updateProj := entities.Project{
		ID:   req.ID,
		Flow: req.Flow,
	}

	if err := h.projSrv.UpdateProject(updateProj, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project flow", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project flow updated successfully"})
}

// UpdateProjectTypeConfig updates the project type configuration
func (h *ProjectHandler) UpdateProjectTypeConfig(c *gin.Context) {
	var req request.UpdateProjectTypeConfig
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	// Validate required fields
	if req.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
		return
	}

	userID := c.GetHeader("X-User-Id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in headers"})
		return
	}

	updateProj := entities.Project{
		ID:          req.ID,
		TypesConfig: req.TypesConfig,
	}

	if err := h.projSrv.UpdateProject(updateProj, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project type configuration", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project type configuration updated successfully"})
}

// DeleteProject deletes a project
func (h *ProjectHandler) DeleteProject(c *gin.Context) {
	userId := c.Query("userId")
	projectId := c.Query("projectId")

	if userId == "" || projectId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID and Project ID are required"})
		return
	}

	if err := h.projSrv.DeleteProject(userId, projectId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}

// PublicShare marks a project as public.
func (h *ProjectHandler) PublicShare(c *gin.Context) {
	projectID := c.Param("projectId")
	if projectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
		return
	}

	err := h.projSrv.PublicShare(projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to make project public", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project is now public"})
}

func (h *ProjectHandler) DepublicShare(c *gin.Context) {
	projectID := c.Param("projectId")
	if projectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
		return
	}

	err := h.projSrv.DepublicShare(projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to make project private", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project is now private"})
}

// GetPublicSharedProjects retrieves all publicly shared projects.
func (h *ProjectHandler) GetPublicSharedProjects(c *gin.Context) {
	projects, err := h.projSrv.GetPublicSharedProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve public projects", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"projects": projects})
}

func (h *ProjectHandler) CanCloneProject(c *gin.Context) {
	// Retrieve projectId from URL params
	projectId := c.Param("projectId")
	if projectId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
		return
	}

	// Call the service method to check if the project can be cloned
	canClone, err := h.projSrv.CanCloneProject(projectId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check clone availability", "details": err.Error()})
		return
	}

	// Return the result
	c.JSON(http.StatusOK, gin.H{"canClone": canClone})
}

// CloneProject handles the cloning of a project by a user.
func (h *ProjectHandler) CloneProject(c *gin.Context) {
	// Get projectId from the URL params
	projectId := c.Param("projectId")
	if projectId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
		return
	}

	// Get newOwnerId from the request body or query params
	var req struct {
		NewOwnerId string `json:"newOwnerId" binding:"required"`
	}

	// Bind the request body to extract newOwnerId
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	// Check if newOwnerId is provided
	if req.NewOwnerId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "New owner ID is required"})
		return
	}

	// Call the service method to clone the project
	clonedProject, err := h.projSrv.CloneProject(projectId, req.NewOwnerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clone project", "details": err.Error()})
		return
	}

	// Return the cloned project details in the response
	c.JSON(http.StatusOK, gin.H{"message": "Project cloned successfully", "project": clonedProject})
}
