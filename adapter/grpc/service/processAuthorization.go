package service

import (
	"context"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
	"github.com/renatospaka/payment-transaction/core/usecase"
)

type AuthorizationService struct {
	usecases *usecase.TransactionUsecase
	service  *pb.AuthorizationServiceClient
}

func NewAuthorizationService(usecases *usecase.TransactionUsecase, service *pb.AuthorizationServiceClient) *AuthorizationService {
	return &AuthorizationService{
		usecases: usecases,
		service:  service,
	}
}

func (a *AuthorizationService)  Process(ctx context.Context, in *pb.AuthorizationRequest) (*pb.AuthorizationResponse, error) {
	// response, err := a.service.

	return nil, nil
}
