package service

import (
	"context"
	"log"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
)

func (t *TransactionService) authorizeTransaction(ctx context.Context, in *pb.AuthorizationProcessRequest) (*pb.AuthorizationProcessResponse, error) {
	log.Println("service.transactions.authorizeTransaction")
	log.Printf("SERVICE => in.TransactionId: %s | in.ClientId: %s | in.Value: %f\n", in.TransactionId, in.ClientId, in.Value)

	response, err := t.grpServices.Client.Process(ctx, in)
	authResponse := &pb.AuthorizationProcessResponse{
		AuthorizationId: response.TransactionId,
		ClientId:        response.ClientId,
		TransactionId:   response.TransactionId,
		Status:          response.Status,
		Value:           response.Value,
	}

	if err != nil {
		authResponse.ErrorMessage = response.ErrorMessage
		return authResponse, err
	}
	return authResponse, nil
}
