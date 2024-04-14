package main

import (
	"fmt"
	"grpc-service/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Host
var port int = 8082

type CommonServiceServer struct {
	proto.CommonServiceServer
}

/*실험환경에서는 Insecure SSL을 통해 진행을 한다는점 알아두시기 바랍니다.*/
func main() {
	// Error
	var err error
	// Server listener
	var listener net.Listener
	// Common Service Server struct address
	commonServices := new(CommonServiceServer)

	var host = fmt.Sprintf(":%v", port)

	if listener, err = net.Listen("tcp", host); err != nil {
		log.Fatalf("Fail to initialize listener: %v\n", err)
	}

	log.Printf("Listening on port %v\n", port)

	// Create empty gRPC servers + Register common services server to grpc server
	server := grpc.NewServer()
	proto.RegisterCommonServiceServer(server, commonServices)

	if err = server.Serve(listener); err != nil {
		log.Fatalf("Error while start gRPC server: %v", err)
	}
}
