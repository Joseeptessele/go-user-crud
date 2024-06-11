package main

import (
	"context"
	"log"

	"github.com/Joseeptessele/go-user-crud/src/configuration/database/mongodb"
	"github.com/Joseeptessele/go-user-crud/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("error trying to connect to database, error=%s", err.Error())
	}
	// init dependencies
	controller := initDependencies(database)
	
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, controller)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
