package service

import (
	"context"
	"errors"
	"fmt"
	"head-api/proto"
	"io"
	"log"
	"net/http"
	"os"

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
	conn, err = grpc.Dial(os.Getenv("GRPC_REQUEST_ENDPOINT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
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
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Fail while processing: %v", err)})
			return
		}
		results = append(results, int(res.ResponseNumber))
	}

	ctx.JSON(http.StatusOK, gin.H{"datas": results})
}

func ClientStreamHandler(ctx *gin.Context) {
	var client *ServiceClient
	// Error
	var err error

	// Query string
	var qs = new(CommonQuery)

	// Bind querystring to struct
	if err = ctx.ShouldBindQuery(qs); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate connection
	client, err = GetGrpcConnection()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Fail to connect gRPC endpoint: %v", err)})
		return
	}
	// Gurantee that connection will be closed
	defer client.Connection.Close()
	client.GenerateClient()

	stream, clientStreamErr := client.Client.ClientStreamingCommunication(context.Background())
	if clientStreamErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Fail to initialize stream: %v", err)})
	}

	// Send numbers to server as stream
	for i := qs.From; i < qs.To; i++ {
		stream.Send(&proto.CommonRequest{
			To: i,
		})
	}

	// End stream and receive result
	streamResponse, serverResponseError := stream.CloseAndRecv()
	if serverResponseError != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Fail to get response from server: %v", err)})
	}
	ctx.JSON(http.StatusOK, gin.H{"datas": streamResponse.ResponseNumbers})
}

func ServerStreamHandler(ctx *gin.Context) {
	var client *ServiceClient
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

	// Generate connection
	client, err = GetGrpcConnection()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Fail to connect gRPC endpoint: %v", err)})
		return
	}
	// Gurantee that connection will be closed
	defer client.Connection.Close()
	client.GenerateClient()
	stream, err := client.Client.ServerStreamingCommunication(context.Background(), &proto.ServerStreamRequest{
		From: qs.From,
		To:   qs.To,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Fail while receiving value from server: %v", err)})
		return
	}
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Fail while receiving value from server: %v", err)})
			return
		}
		results = append(results, int(recv.ResponseNumber))
	}
	ctx.JSON(http.StatusOK, gin.H{"datas": results})
}

func BiDirectionalStreamHandler(ctx *gin.Context) {
	var client *ServiceClient
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

	// Generate connection
	client, err = GetGrpcConnection()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Fail to connect gRPC endpoint: %v", err)})
		return
	}

	// Gurantee that connection will be closed
	defer client.Connection.Close()
	client.GenerateClient()

	// Get stream
	stream, streamErr := client.Client.BiDirectionalCommunication(context.Background())
	if streamErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error while generating stream: %v", streamErr)})
		return
	}

	// Generate channel to block context
	blockingChannel := make(chan struct{})

	// Client -> Server Stream
	go func() {
		var clientStreamError error
		for i := qs.From; i < qs.To; i++ {
			clientStreamError = stream.Send(&proto.CommonRequest{
				To: i,
			})
			if clientStreamError != nil {
				log.Printf("Error while sending data: %v", clientStreamError)
			}
		}

		// Close client side stream
		stream.CloseSend()
	}()

	// Server -> Client Stream
	go func() {
		var recv *proto.CommonResponse
		var serverStreamError error
		for {
			recv, serverStreamError = stream.Recv()

			// If server stream end
			if serverStreamError == io.EOF {
				break
			}
			if serverStreamError != nil {
				log.Printf("Error while receiving message from server: %v", serverStreamError)
			}
			results = append(results, int(recv.ResponseNumber))
		}
		close(blockingChannel)
	}()
	// Block untile receiving channel
	<-blockingChannel
	ctx.JSON(http.StatusOK, gin.H{"datas": results})
}
