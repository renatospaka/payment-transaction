package main

import (
	"context"
	"log"
	"net/http"

	httpServer "github.com/renatospaka/payment-transaction/adapter/httpServer"
	repository "github.com/renatospaka/payment-transaction/adapter/postgres"
	"github.com/renatospaka/payment-transaction/adapter/rest/controller"
	"github.com/renatospaka/payment-transaction/utils/configs"
)

func main() {
	log.Println("iniciando a aplicação")
	configs, err := configs.LoadConfig(".")
	if err != nil {
		log.Panic(err)
	}

	//open connection to the database
	ctx := context.Background()
	repository.NewPostgresDatabase()
	controllers := controller.NewTransactionController()
	handler := httpServer.NewHttpServer(ctx, controllers)

	//start web server
	log.Println("servidor escutando porta:", configs.WEBServerPort)
	http.ListenAndServe(":"+configs.WEBServerPort, handler.Server)
}
