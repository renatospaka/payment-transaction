package usecase

import (
	"context"
	"log"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
	"github.com/renatospaka/payment-transaction/core/dto"
	"github.com/renatospaka/payment-transaction/core/entity"
	"github.com/renatospaka/payment-transaction/core/service"
)

func (t *TransactionUsecase) reprocessTransactionPendingAuthorization(ctx context.Context, tr *dto.TransactionCreateDto) (*dto.TransactionDto, error) {
	log.Println("usecase.transactions.reprocessTransactionPendingAuthorization")

	if t.services == nil {
		return nil, service.ErrGRPCNotDefined
	}
	if tr.ID == "" {
		return nil, entity.ErrTransactionIDIsRequired
	}

	// find the transaction by transaction id
	transaction, err := t.repo.Find(tr.ID)
	if err != nil {
		return nil, err
	}

	if transaction.GetStatus() != entity.TR_PENDING {
		return nil, entity.ErrCannotReprocess
	}

	id := transaction.GetID()
	clientId := transaction.GetClientID()
	value := transaction.GetValue()
	tr.ID = id

	// Call the Reprocess Authorization service and update status accordingly
	auth := &pb.AuthorizationReprocessRequest{
		ClientId:      clientId,
		TransactionId: id,
		Value:         value,
	}
	authorize, err := t.services.ReprocessTransactionPendingAuthorization(ctx, auth)
	if err != nil {
		return nil, err
	}

	// request the gRPC server to authorize the transaction
	tran := &dto.TransactionAuthorizeDto{
		ID:              id,
		ClientID:        clientId,
		AuthorizationId: authorize.AuthorizationId,
		Value:           value,
	}
	if authorize.Status == entity.TR_APPROVED {
		err = t.approveTransaction(ctx, tran)
	} else if authorize.Status == entity.TR_DENIED {
		err = t.denyTransaction(ctx, tran)
	}
	if err != nil {
		return nil, err
	}

	// Find the new Transaction the reply to the caller
	created, err := t.findTransactionById(ctx, id)
	if err != nil {
		return nil, err
	}
	return created, nil
}
