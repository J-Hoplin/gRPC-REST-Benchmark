/*
사용자 프런트 API는 라우트 관리를 위해

Golang Gin Gonic 으로 작성했습니다.
*/
package main

import "github.com/gin-gonic/gin"

func main() {

	// Using default to make logger and recovery middleware attached to application
	router := gin.Default()

	// Enroll Router
	RestHandler(router)
	GrpcHandler(router)
}
