package service

import (
	"context"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
)

type AuthorizationServiceInterface interface {
	AuthorizeNewTransaction(context.Context, *pb.AuthorizationProcessNewRequest) (*pb.AuthorizationProcessNewResponse, error)
	ReprocessTransactionPendingAuthorization(context.Context, *pb.AuthorizationReprocessPendingRequest) (*pb.AuthorizationReprocessPendingResponse, error)
}
