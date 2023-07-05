package service

// import (
// 	"context"

// 	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
// )

// type AuthorizationService struct {
// }

// func NewAuthorizationService() *AuthorizationService {
// 	return &AuthorizationService{}
// }

// func (a *AuthorizationService) Process(ctx context.Context, in *pb.AuthorizationRequest) (*pb.AuthorizationResponse, error) {
// 	auth := &pb.AuthorizationRequest{
// 		ClientId:      in.ClientId,
// 		TransactionId: in.TransactionId,
// 		Value:         in.Value,
// 	}

// 	// response, err := a.usecases.Authorize(auth)
// 	response, err := interface{}, interface{}
// 	authResponse := &pb.AuthorizationResponse{
// 		AuthorizationId: response.TransactionId,
// 		ClientId:        response.ClientId,
// 		TransactionId:   response.TransactionId,
// 		Status:          response.Status,
// 		Value:           response.Value,
// 	}

// 	if err != nil {
// 		authResponse.ErrorMessage = response.ErrorMessage
// 		return authResponse, nil
// 	}
// 	return authResponse, nil
// }
