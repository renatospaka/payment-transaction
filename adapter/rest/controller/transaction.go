package controller

import (
	"net/http"

	"github.com/renatospaka/payment-transaction/core/usecase"
)

type TransactionController struct {
	usecases *usecase.TransactionUsecase
}

func NewTransactionController(usecases *usecase.TransactionUsecase) *TransactionController {
	return &TransactionController{
		usecases: usecases,
	}
}

func (c *TransactionController) Process(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (c *TransactionController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (c *TransactionController) Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (c *TransactionController) Modify(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (c *TransactionController) Remove(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
