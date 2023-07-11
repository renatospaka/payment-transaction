package service

import (
	"context"

	"github.com/renatospaka/payment-transaction/adapter/grpc/client"
	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
)

type TransactionService struct {
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

// Call the Process gRPC service requesting to reprocess the auhorization for this pending transaction
func (t *TransactionService) ReprocessTransactionPendingAuthorization(ctx context.Context, in *pb.AuthorizationReprocessPendingRequest) (*pb.AuthorizationReprocessPendingResponse, error) {
	return t.reprocessTransactionPendingAuthorization(ctx, in)
}
