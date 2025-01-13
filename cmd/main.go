package main

import (
	"log"
	"os"
	graph "project-management-service/external/handler/adaptors/graphql"
	handler "project-management-service/external/handler/adaptors/rest/api"
	"project-management-service/external/handler/router"
	"project-management-service/external/receiver/adaptors/gRPC"
	receiver "project-management-service/external/receiver/adaptors/gRPC/controller"
	gorm "project-management-service/external/repository/adaptors/postgres"
	repository "project-management-service/external/repository/adaptors/postgres/controller"
	"project-management-service/internal/core/service"
	"project-management-service/pkg/db"
	mygrpc "project-management-service/pkg/gRPC"

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
	clientDB := postgresDB.GetClient()

	gorm.SyncDB(clientDB)

	projectRep := repository.NewProjectRepositoryPQ(clientDB)
	projectSrv := service.NewProjectService(projectRep)

	cilent, err := mygrpc.NewGRPCClient(os.Getenv("USER_INFO_SERVICE_HOST"), os.Getenv("USER_INFO_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}

	userGRPCclient := gRPC.NewUserServiceClient(cilent)
	userRev := receiver.NewUserReceiver(userGRPCclient)
	userSrv := service.NewUserService(userRev)

	projectHandler := handler.NewProjectHandler(projectSrv)
	router.RegisterProjectRoutes(r, projectHandler)

	resolver := &graph.Resolver{
		ProjSrv: projectSrv,
		UserSrv: userSrv,
	}

	srv := gqlHandler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
	router.RegisterGQLRoutes(r, srv)

	r.Run(":8081")
}
