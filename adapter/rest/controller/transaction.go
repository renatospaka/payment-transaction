package controller

import "net/http"

type TransactionController struct {
}

func NewTransactionController() *TransactionController {
	return &TransactionController{}
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
