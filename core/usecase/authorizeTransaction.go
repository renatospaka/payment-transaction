package usecase

import (
	"context"
	"log"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
)

func (t *TransactionUsecase) authorizeTransaction(ctx context.Context, in *pb.AuthorizationRequest) (*pb.AuthorizationResponse, error) {
	log.Println("usecase.transactions.authorize")
	
	// Execute the gRPC call
	response := &pb.AuthorizationResponse{}
	auth, err := t.services.AuthorizeTransaction(ctx, in)
	if err != nil {
		response = &pb.AuthorizationResponse{
			AuthorizationId: "",
			ClientId:        in.ClientId,
			TransactionId:   in.TransactionId,
			Status:          auth.Status,
			Value:           in.Value,
			ErrorMessage:    auth.ErrorMessage,
		}
		return response, err
	}

	response = &pb.AuthorizationResponse{
		AuthorizationId: auth.AuthorizationId,
		ClientId:        in.ClientId,
		TransactionId:   in.TransactionId,
		Status:          auth.Status,
		Value:           in.Value,
	}
	return response, nil
}
