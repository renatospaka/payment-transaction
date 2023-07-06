package service

import (
	"context"
	"log"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
)

func (t *TransactionService) authorizeTransaction(ctx context.Context, in *pb.AuthorizationRequest) (*pb.AuthorizationResponse, error) {
	log.Println("service.transactions.authorizeTransaction")
	// auth := &pb.AuthorizationRequest{
	// 	ClientId:      in.ClientId,
	// 	TransactionId: in.TransactionId,
	// 	Value:         in.Value,
	// }

	// response, err := t.usecases.Authorize(auth)
	// authResponse := &pb.AuthorizationResponse{
	// 	AuthorizationId: response.TransactionId,
	// 	ClientId:        response.ClientId,
	// 	TransactionId:   response.TransactionId,
	// 	Status:          response.Status,
	// 	Value:           response.Value,
	// }
	authResponse := &pb.AuthorizationResponse{
		AuthorizationId: "",
		ClientId:        "",
		TransactionId:   "",
		Status:          "",
		Value:           0,
		ErrorMessage:    "",
	}

	// if err != nil {
	// 	authResponse.ErrorMessage = response.ErrorMessage
	// 	return authResponse, nil
	// }
	return authResponse, nil
}
