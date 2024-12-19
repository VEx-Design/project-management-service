package main

import (
	"log"
	handler "project-management-service/external/handler/adaptors/gin/api"
	"project-management-service/external/handler/adaptors/gin/router"
	repository "project-management-service/external/repository/adaptors/mongodb/controller"
	"project-management-service/internal/core/service"
	"project-management-service/pkg/db"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	mongoDB := db.ConnectToMongo()
	defer mongoDB.Disconnect()

	projectRep := repository.NewProjectRepositoryMongo(mongoDB.GetClient())
	projectSrv := service.NewProjectService(projectRep)
	projectHandler := handler.NewProjectHandler(projectSrv)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://your-frontend-domain.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.RegisterProjectRoutes(r, projectHandler)
	r.Run(":8080")
}
