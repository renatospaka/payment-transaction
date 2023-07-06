package service

import (
	"context"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
)

type AuthorizationServiceInterface interface {
	AuthorizeTransaction(context.Context, *pb.AuthorizationRequest) (*pb.AuthorizationResponse, error)
}
