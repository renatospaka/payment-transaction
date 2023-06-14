package main

import (
	"context"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	server "github.com/renatospaka/transaction/adapter/http/rest"
	"github.com/renatospaka/transaction/adapter/http/rest/controller"
	repository "github.com/renatospaka/transaction/adapter/postgres"
	"github.com/renatospaka/transaction/utils/configs"
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
	handler := server.NewHttpServer(ctx, controllers)
	
	//start web server
	log.Println("servidor escutando porta:", configs.DBPort)
	http.ListenAndServe(":" + configs.DBPort, handler.Server)
}
