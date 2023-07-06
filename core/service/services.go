package service

import (
	"context"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
)

type AuthorizationService struct {}

func NewAuthorizationService() *AuthorizationService {
	return &AuthorizationService{}
}

type TransactionService struct {}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}


// Call the Process gRPC service requesting auhorization for this transaction
func (t *TransactionService) AuthorizeTransaction(ctx context.Context, in *pb.AuthorizationRequest) (*pb.AuthorizationResponse, error) {
	return t.authorizeTransaction(ctx, in)
}
