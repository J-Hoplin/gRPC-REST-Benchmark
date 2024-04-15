package service

import (
	"context"
	"grpc-service/proto"
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
