package main

import (
	"go-jwt-implement/controller"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	server := gin.New()
	server.GET("info", controller.Test)

	server.Run(":8080")
}
