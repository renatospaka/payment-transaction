package service

import (
	"context"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
	"github.com/renatospaka/payment-transaction/adapter/grpc/client"
)

type TransactionService struct {	
	// pb.UnimplementedAuthorizationServiceServer
	grpServices *client.GrpcClient
}

func NewTransactionService(services *client.GrpcClient) *TransactionService {
	return &TransactionService{
		grpServices: services,
	}
}

// Call the Process gRPC service requesting auhorization for this transaction
func (t *TransactionService) AuthorizeTransaction(ctx context.Context, in *pb.AuthorizationProcessRequest) (*pb.AuthorizationProcessResponse, error) {
	return t.authorizeTransaction(ctx, in)
}
