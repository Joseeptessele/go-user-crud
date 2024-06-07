package main

import (
	"log"

	"github.com/Joseeptessele/go-user-crud/src/configuration/database/mongodb"
	"github.com/Joseeptessele/go-user-crud/src/controller"
	"github.com/Joseeptessele/go-user-crud/src/controller/routes"
	"github.com/Joseeptessele/go-user-crud/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	mongodb.InitConnection()

	// init dependencies
	service := service.NewUserDomainService()
	controller := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, controller)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
