package service

import (
	"context"
	"grpc-service/proto"
	"io"
	"log"
)

type CommonServiceServer struct {
	proto.CommonServiceServer
}

// Implement Unary Communication
func (s *CommonServiceServer) UnaryCommunication(ctx context.Context, input *proto.CommonRequest) (*proto.CommonResponse, error) {
	to := input.To
	var sumAll, i int64
	var response = new(proto.CommonResponse)
	for ; i < to; i++ {
		sumAll += i * i
	}
	response.ResponseNumber = sumAll

	return response, nil
}

// Implement Client Streaming communication
func (s *CommonServiceServer) ClientStreamingCommunication(stream proto.CommonService_ClientStreamingCommunicationServer) error {
	// Error
	var err error
	// Received
	var recv *proto.CommonRequest

	// Array for response
	var responseNumbers = make([]int64, 0)
	var sumAll int64
	for {
		// Recv는 내부적으로 다음 값이 올때까지 blocking state를 유지합니다.
		recv, err = stream.Recv()

		// Stream 측에서 io.EOF(End Of File)에러를 보내면, 이는 스트림이 끝났음을 의미합니다.
		if err == io.EOF {
			break
		}

		// 다른 오류인 경우
		if err != nil {
			log.Fatalf("Fail to process received value from client: %v\n", err)
		}

		// Initialize sumAll to 0
		sumAll = 0
		for i := 0; i < int(recv.To); i++ {
			sumAll += int64(i * i)
		}

		responseNumbers = append(responseNumbers, sumAll)
	}

	// Send message to client and close network connection
	stream.SendAndClose(&proto.ClientStreamResponse{
		ResponseNumbers: responseNumbers,
	})
	return nil
}

// Implement Server stream communication
func (s *CommonServiceServer) ServerStreamingCommunication(input *proto.ServerStreamRequest, stream proto.CommonService_ServerStreamingCommunicationServer) error {
	// Get arrays from client
	var tos = input.Tos
	// Error
	var err error

	for _, val := range tos {
		var sumAll = 0
		for i := 0; i < int(val); i++ {
			sumAll += i * i
		}
		err = stream.Send(&proto.CommonResponse{
			ResponseNumber: int64(sumAll),
		})
		if err != nil {
			log.Fatalf("Error while sending message to client through stream: %v\n", err)
		}
	}

	// End stream
	return nil
}

func (s *CommonServiceServer) BiDirectionalCommunication(stream proto.CommonService_BiDirectionalCommunicationServer) error {
	// Error
	var err error

	// Received Type
	var recv *proto.CommonRequest

	for {
		recv, err = stream.Recv()

		// If client end stream
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while bi-directional stream: %v\n", err)
		}
		var sumAll = 0
		for i := 0; i < int(recv.To); i++ {
			sumAll += i * i
		}
		err = stream.Send(&proto.CommonResponse{
			ResponseNumber: int64(sumAll),
		})

		if err != nil {
			log.Fatalf("Error while sending response to stream: %v\n", err)
		}
	}
}
