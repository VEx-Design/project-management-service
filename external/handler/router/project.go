package router

import (
	"context"
	handler "project-management-service/external/handler/adaptors/rest/api"

	gqlHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/gin-gonic/gin"
)

type requestID string

const RequestIDKey = requestID("user_id")

func RegisterProjectRoutes(router *gin.Engine, projHandler *handler.ProjectHandler) {
	router.POST("/project", projHandler.CreateProject)
	router.GET("/project", projHandler.GetMyProject)
	router.PUT("/project", projHandler.UpdateProject)
	router.PUT("/project/flow", projHandler.UpdateProjectFlow)
	router.PUT("/project/type", projHandler.UpdateProjectTypeConfig)
	router.DELETE("/project", projHandler.DeleteProject)

	router.PUT("/:projectId/public", projHandler.PublicShare)     // Make project public
	router.PUT("/:projectId/depublic", projHandler.DepublicShare) // Make project private
	router.GET("/public", projHandler.GetPublicSharedProjects)    // Get public projects
	router.PUT("/clone/check/:projectId", projHandler.CanCloneProject)
	router.POST("/clone/:projectId", projHandler.CloneProject)
}

func RegisterGQLRoutes(router *gin.Engine, srv *gqlHandler.Server) {
	// GraphQL Playground route
	router.GET("/playground", func(c *gin.Context) {
		userID := c.GetHeader("X-User-Id")

		ctx := context.WithValue(c.Request.Context(), RequestIDKey, userID)

		c.Request = c.Request.WithContext(ctx)
		playground.Handler("GraphQL Playground", "/query").ServeHTTP(c.Writer, c.Request)
	})

	// GraphQL API route
	router.POST("/query", func(c *gin.Context) {
		userID := c.GetHeader("X-User-Id")

		ctx := context.WithValue(c.Request.Context(), RequestIDKey, userID)

		c.Request = c.Request.WithContext(ctx)
		srv.ServeHTTP(c.Writer, c.Request)
	})

}
