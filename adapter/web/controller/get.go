package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

// Return a specific Transaction
func (c *TransactionController) Get(w http.ResponseWriter, r *http.Request) {
	log.Println("http.transactions.get")

	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tr, err := c.usecases.FindTransactionById(id)
	if tr == nil || err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("error: " + err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&tr)
}
