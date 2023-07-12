package usecase

import (
	"context"
	"log"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
)

func (t *TransactionUsecase) authorizeNewTransaction(ctx context.Context, in *pb.AuthorizationProcessNewRequest) (*pb.AuthorizationProcessNewResponse, error) {
	log.Println("usecase.transactions.authorizeNewTransaction")

	// Execute the gRPC call
	response := &pb.AuthorizationProcessNewResponse{}
	auth, err := t.services.AuthorizeNewTransaction(ctx, in)
	if err != nil {
		response = &pb.AuthorizationProcessNewResponse{
			AuthorizationId: "",
			ClientId:        in.ClientId,
			TransactionId:   in.TransactionId,
			Status:          auth.Status,
			Value:           in.Value,
			ErrorMessage:    auth.ErrorMessage,
		}
		return response, err
	}

	response = &pb.AuthorizationProcessNewResponse{
		AuthorizationId: auth.AuthorizationId,
		ClientId:        in.ClientId,
		TransactionId:   in.TransactionId,
		Status:          auth.Status,
		Value:           in.Value,
	}
	return response, nil
}
