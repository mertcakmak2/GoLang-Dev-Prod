package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()
	router.GET("/hello", hello)

	server_port := fmt.Sprintf("%s%s", ":", getEnv("SERVER_PORT"))
	router.Run(server_port)
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello from: "+getEnv("MODE"))
}

func getEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dotenv := os.Getenv(key)
	return dotenv
}
