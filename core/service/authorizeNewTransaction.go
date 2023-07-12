package service

import (
	"context"
	"log"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
)

func (t *TransactionService) authorizeNewTransaction(ctx context.Context, in *pb.AuthorizationProcessNewRequest) (*pb.AuthorizationProcessNewResponse, error) {
	log.Println("service.transactions.authorizeNewTransaction")

	response, err := t.grpServices.Client.ProcessNewAuthorization(ctx, in)
	authResponse := &pb.AuthorizationProcessNewResponse{
		AuthorizationId: response.TransactionId,
		ClientId:        response.ClientId,
		TransactionId:   response.TransactionId,
		Status:          response.Status,
		Value:           response.Value,
		ErrorMessage:    "",
	}

	if err != nil {
		authResponse.ErrorMessage = response.ErrorMessage
		return authResponse, err
	}
	return authResponse, nil
}
