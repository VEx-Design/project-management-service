package main

import (
	"log"
	graph "project-management-service/external/handler/adaptors/graphql"
	handler "project-management-service/external/handler/adaptors/rest/api"
	"project-management-service/external/handler/router"
	gorm "project-management-service/external/repository/adaptors/postgres"
	repository "project-management-service/external/repository/adaptors/postgres/controller"
	"project-management-service/internal/core/service"
	"project-management-service/pkg/db"

	gqlHandler "github.com/99designs/gqlgen/graphql/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	r := gin.Default()

	postgresDB := db.ConnectToPG()
	client := postgresDB.GetClient()

	gorm.SyncDB(client)

	projectRep := repository.NewProjectRepositoryPQ(client)
	projectSrv := service.NewProjectService(projectRep)
	projectHandler := handler.NewProjectHandler(projectSrv)
	router.RegisterProjectRoutes(r, projectHandler)

	resolver := &graph.Resolver{
		ProjSrv: projectSrv,
	}

	srv := gqlHandler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
	router.RegisterGQLRoutes(r, srv)

	r.Run(":8081")
}
