package router

import (
	handler "project-management-service/external/handler/adaptors/rest/api"

	gqlHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/gin-gonic/gin"
)

func RegisterProjectRoutes(router *gin.Engine, projHandler *handler.ProjectHandler) {
	// REST API routes
	api := router.Group("/api/v1")
	{
		api.POST("/project", projHandler.CreateProject)
		api.GET("/project", projHandler.GetMyProject)
	}
}

func RegisterGQLRoutes(router *gin.Engine, srv *gqlHandler.Server) {
	// GraphQL Playground route
	router.GET("/playground", func(c *gin.Context) {
		playground.Handler("GraphQL Playground", "/query").ServeHTTP(c.Writer, c.Request)
	})

	// GraphQL API route
	router.POST("/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})
}
