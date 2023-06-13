package main

import (
	"context"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	server "github.com/renatospaka/transaction/adapter/http/rest"
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
	handler := server.NewHttpServer(ctx)
	
	//start web server
	log.Println("servidor escutando porta:", configs.DBPort)
	http.ListenAndServe(":" + configs.DBPort, handler.Server)
}
