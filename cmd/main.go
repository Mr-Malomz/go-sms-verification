package main

import (
	"go-sms-verification/api"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Printf("Starting broker service on port 80")
	router := gin.Default()

	//initialize config
	app := api.Config{Router: router}

	//routes
	app.Routes()

	router.Run(":80")
}
