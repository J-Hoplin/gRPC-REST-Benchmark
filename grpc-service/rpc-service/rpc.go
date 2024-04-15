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
func UnaryCommunication(ctx context.Context, input *proto.CommonRequest) (*proto.CommonResponse, error) {
	to := input.To
	var sumAll, i int64
	var response = new(proto.CommonResponse)
	for ; i < to; i++ {
		sumAll += i * i
	}
	response.ResponseNumber = sumAll

	return response, nil
}

func ClientStreamingCommunication(stream proto.CommonService_ClientStreamingCommunicationServer) error {
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
		var i int64 = 0
		for ; i < recv.To; i++ {
			sumAll += i * i
		}

		responseNumbers = append(responseNumbers, sumAll)
	}

	// Send message to client and close network connection
	stream.SendAndClose(&proto.ClientStreamResponse{
		ResponseNumbers: responseNumbers,
	})
	return nil
}
