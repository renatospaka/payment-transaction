package usecase

import (
	"context"
	"log"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
)

func (t *TransactionUsecase) authorizeTransaction(ctx context.Context, in *pb.AuthorizationProcessRequest) (*pb.AuthorizationProcessResponse, error) {
	log.Println("usecase.transactions.authorizeTransaction")
	log.Printf("USECASE => in.TransactionId: %s | in.ClientId: %s | in.Value: %d\n", in.TransactionId, in.ClientId, in.Value)

	// Execute the gRPC call
	response := &pb.AuthorizationProcessResponse{}
	auth, err := t.services.AuthorizeTransaction(ctx, in)
	if err != nil {
		response = &pb.AuthorizationProcessResponse{
			AuthorizationId: "",
			ClientId:        in.ClientId,
			TransactionId:   in.TransactionId,
			Status:          auth.Status,
			Value:           in.Value,
			ErrorMessage:    auth.ErrorMessage,
		}
		return response, err
	}

	response = &pb.AuthorizationProcessResponse{
		AuthorizationId: auth.AuthorizationId,
		ClientId:        in.ClientId,
		TransactionId:   in.TransactionId,
		Status:          auth.Status,
		Value:           in.Value,
	}
	return response, nil
}
