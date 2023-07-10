package service

import (
	"context"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
)

type AuthorizationServiceInterface interface {
	AuthorizeTransaction(context.Context, *pb.AuthorizationProcessRequest) (*pb.AuthorizationProcessResponse, error)
	ReprocessTransactionPendingAuthorization(context.Context, *pb.AuthorizationReprocessRequest) (*pb.AuthorizationReprocessResponse, error)
}
