package handler

import (
	"head-api/service"

	"github.com/gin-gonic/gin"
)

func EnrollRestHandler(router *gin.Engine) {
	group := router.Group("rest")
	{
		group.GET("", service.RestHandler)
	}
}
