package usecase

import (
	"context"
	"log"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
	"github.com/renatospaka/payment-transaction/core/dto"
	"github.com/renatospaka/payment-transaction/core/repository"
	"github.com/renatospaka/payment-transaction/core/service"
)

type TransactionUsecase struct {
	repo     repository.TransactionInterface
	services service.AuthorizationServiceInterface
}

var ctx context.Context

func NewTransactionUsecase(repo repository.TransactionInterface) *TransactionUsecase {
	ctx = context.Background()
	return &TransactionUsecase{
		repo: repo,
	}
}

// Inject the gRPC service into the usecases
func (t *TransactionUsecase) SetServices(services service.AuthorizationServiceInterface) {
	t.services = services
}

// Create a new transaction and process its authorization
func (t *TransactionUsecase) CreateTransactionAndProcessAuthorization(tr *dto.TransactionCreateDto) (*dto.TransactionDto, error) {
	return t.createTransactionAndProcessAuthorization(ctx, tr)
}

// Find an existing transaction by its id
func (t *TransactionUsecase) FindTransactionById(id string) (*dto.TransactionDto, error) {
	return t.findTransactionById(ctx, id)
}

// There is no business validations related to retrieving all transactions at the usecase level
func (t *TransactionUsecase) FindAll(page int, limit int) (dto.TransactionFindAllResponseDto, error) {
	return t.findAllTransactions(ctx, page, limit)
}

// All business validations related to deleting transactions occur at the usecase level
func (t *TransactionUsecase) Delete(id string) error {
	log.Println("usecase.transactions.delete")
	return t.deleteTransaction(ctx, id)
}

// All business validations related to updating transactions occur at the usecase level
func (t *TransactionUsecase) Update(tr *dto.TransactionUpdateDto) error {
	return t.updateTransaction(ctx, tr)
}

// Request to the gRPC Server to process the authorization
func (t *TransactionUsecase) Authorize(auth *pb.AuthorizationProcessRequest) (*pb.AuthorizationProcessResponse, error) {
	return t.authorizeTransaction(ctx, auth)
}
