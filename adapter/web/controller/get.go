package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/renatospaka/payment-transaction/utils/configs"
)

// Return a specific Transaction
func (c *TransactionController) Get(w http.ResponseWriter, r *http.Request) {
	log.Println("http.transactions.get")

	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cng, _ := configs.LoadConfig("../../app/")
	ctx := r.Context()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Duration(cng.WEBServerTimeOut) * time.Millisecond))
	defer cancel()

	c.usecases.SetContext(ctx)
	time.Sleep(500 * time.Millisecond)

	tr, err := c.usecases.FindTransactionById(id)
	if tr == nil || err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("error: " + err.Error())
		return
	}

	select {
	case <-ctx.Done():
		w.WriteHeader(http.StatusRequestTimeout)
		return
	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&tr)
	}
}
