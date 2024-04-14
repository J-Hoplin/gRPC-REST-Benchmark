package handler

import (
	"github.com/gin-gonic/gin"
)

func EnrollGrpcHandler(router *gin.Engine) {
	group := router.Group("grpc")
	{
		group.GET("unary", func(ctx *gin.Context) {})
		subGroup := group.Group("stream")
		{
			// Client Stream
			subGroup.GET("client", func(ctx *gin.Context) {})

			// Server Stream
			subGroup.GET("server", func(ctx *gin.Context) {})

			// Bi-directional Stream
			subGroup.GET("bi", func(ctx *gin.Context) {})
		}
	}
}
