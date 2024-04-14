package main

/*
사용자 프런트 API는 라우트 관리를 위해

Golang Gin Gonic 으로 작성했습니다.
*/

import (
	"fmt"
	"head-api/handler"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var host int = 8080

func main() {

	// Load dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Fail to load .env file: %v\n", err)
	}

	// Using default to make logger and recovery middleware attached to application
	router := gin.Default()

	// Enroll Router
	handler.EnrollRestHandler(router)
	handler.EnrollGrpcHandler(router)

	router.Run(fmt.Sprintf(":%d", host))
}
