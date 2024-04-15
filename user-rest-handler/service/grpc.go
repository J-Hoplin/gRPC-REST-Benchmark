package service

import (
	"context"
	"errors"
	"fmt"
	"head-api/proto"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Connection *grpc.ClientConn
	Client     proto.CommonServiceClient
}

func (s *ServiceClient) GenerateClient() (err error) {
	// If connection not initialized
	if s.Connection == nil {
		err = errors.New("CONNECTION_NOT_INITIALIZED")
		return
	}
	s.Client = proto.NewCommonServiceClient(s.Connection)
	return
}

func GetGrpcConnection() (*ServiceClient, error) {
	var err error
	var conn *grpc.ClientConn

	serverClient := new(ServiceClient)

	conn, err = grpc.NewClient(":8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	serverClient.Connection = conn

	return serverClient, err
}

func UnaryHandler(ctx *gin.Context) {
	// Error
	var err error
	// Query string
	var qs = new(CommonQuery)
	// Array list
	var results = []int{}
	// Bind querystring to struct
	if err = ctx.ShouldBindQuery(qs); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Client struct
	var client *ServiceClient
	// Generate connection
	client, err = GetGrpcConnection()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Fail to connect gRPC endpoint: %v", err)})
		return
	}

	// Gurantee that connection will be closed
	defer client.Connection.Close()
	// Generate client with connection
	client.GenerateClient()

	for i := qs.From; i < qs.To; i++ {
		res, err := client.Client.UnaryCommunication(context.Background(), &proto.CommonRequest{
			To: i,
		})
		if err != nil {
			log.Fatalf("Fail while processing: %v", err)
		}
		results = append(results, int(res.ResponseNumber))
	}

	ctx.JSON(http.StatusOK, gin.H{"datas": results})
}

func ClientStreamHandler(ctx *gin.Context) {
	var client *ServiceClient
	var err error

	// Generate connection
	client, err = GetGrpcConnection()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Fail to connect gRPC endpoint: %v", err)})
		return
	}
	// Gurantee that connection will be closed
	defer client.Connection.Close()
	client.GenerateClient()
}

func ServerStreamHandler(ctx *gin.Context) {
	var client *ServiceClient
	var err error

	// Generate connection
	client, err = GetGrpcConnection()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Fail to connect gRPC endpoint: %v", err)})
		return
	}
	// Gurantee that connection will be closed
	defer client.Connection.Close()
	client.GenerateClient()
}

func BiDirectionalStreamHandler(ctx *gin.Context) {
	var client *ServiceClient
	var err error

	// Generate connection
	client, err = GetGrpcConnection()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Fail to connect gRPC endpoint: %v", err)})
		return
	}
	// Gurantee that connection will be closed
	defer client.Connection.Close()
	client.GenerateClient()
}
