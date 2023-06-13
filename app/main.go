package main

import (
	"log"

	_ "github.com/lib/pq"
	repository "github.com/renatospaka/transaction/adapter/postgres"
	// "github.com/renatospaka/transaction/utils/configs"
)

func main() {
	log.Println("iniciando a aplicação")
	// configs, err := configs.LoadConfig(".")
	// if err != nil {
	// 	log.Panic(err)
	// }

	//open connection to the database
	repository.NewPostgresDatabase()

	
}
