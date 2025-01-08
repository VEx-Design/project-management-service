package main

import (
	"log"
	handler "project-management-service/external/handler/adaptors/rest/api"
	"project-management-service/external/handler/router"
	gorm "project-management-service/external/repository/adaptors/postgres"
	repository "project-management-service/external/repository/adaptors/postgres/controller"
	"project-management-service/internal/core/service"
	"project-management-service/pkg/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	postgresDB := db.ConnectToPG()
	client := postgresDB.GetClient()

	gorm.SyncDB(client)

	projectRep := repository.NewProjectRepositoryPQ(client)
	projectSrv := service.NewProjectService(projectRep)
	projectHandler := handler.NewProjectHandler(projectSrv)

	r := gin.Default()
	router.RegisterProjectRoutes(r, projectHandler)
	r.Run(":8081")
}
