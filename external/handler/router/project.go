package router

import (
	handler "project-management-service/external/handler/adaptors/rest/api"

	gqlHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/gin-gonic/gin"
)

func RegisterProjectRoutes(router *gin.Engine, projHandler *handler.ProjectHandler) {
	router.POST("/project", projHandler.CreateProject)
	router.GET("/project", projHandler.GetMyProject)
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
