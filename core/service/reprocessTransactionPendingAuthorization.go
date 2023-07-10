package service

import (
	"context"
	"log"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
)

func (t *TransactionService) reprocessTransactionPendingAuthorization(ctx context.Context, in *pb.AuthorizationReprocessRequest) (*pb.AuthorizationReprocessResponse, error) {
	log.Println("service.transactions.reprocessTransactionPendingAuthorization")

	response, err := t.grpServices.Client.Reprocess(ctx, in)
	reprocessResp := &pb.AuthorizationReprocessResponse{
		AuthorizationId: response.AuthorizationId,
		ClientId:        response.ClientId,
		TransactionId:   response.TransactionId,
		Status:          response.Status,
		Value:           response.Value,
		ErrorMessage:    "",
	}

	if err != nil {
		reprocessResp.ErrorMessage = response.ErrorMessage
		return reprocessResp, err
	}
	return reprocessResp, nil
}