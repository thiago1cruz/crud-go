package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/thiago1cruz/crud-go/src/configurations/logger"
	"github.com/thiago1cruz/crud-go/src/configurations/routes"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	err = router.Run(":8080")

	if err != nil {
		log.Fatal(err)
	}

}
