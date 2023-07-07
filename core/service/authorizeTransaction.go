package service

import (
	"context"
	"log"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
)

func (t *TransactionService) authorizeTransaction(ctx context.Context, in *pb.AuthorizationRequest) (*pb.AuthorizationResponse, error) {
	log.Println("service.transactions.authorizeTransaction")

	response, err := t.grpServices.Client.Process(ctx, in)
	authResponse := &pb.AuthorizationResponse{
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
