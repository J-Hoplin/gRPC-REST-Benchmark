package handler

import (
	"head-api/service"

	"github.com/gin-gonic/gin"
)

func EnrollGrpcHandler(router *gin.Engine) {
	group := router.Group("grpc")
	{
		group.GET("unary", service.UnaryHandler)
		subGroup := group.Group("stream")
		{
			// Client Stream
			subGroup.GET("client", service.ClientStreamHandler)

			// Server Stream
			subGroup.GET("server", service.ServerStreamHandler)

			// Bi-directional Stream
			subGroup.GET("bi", service.BiDirectionalStreamHandler)
		}
	}
}
