package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/renatospaka/transaction/utils/configs"
)

func main() {
	log.Println("iniciando a aplicação")
	configs, err := configs.LoadConfig(".")
	if err != nil {
		log.Panic(err)
	}

	//open connection to the database
	log.Println("iniciando conexão com o banco de dados")
	conn := "postgresql://" + configs.DBUser + ":" + configs.DBPassword + "@" + configs.DBHost + "/" + configs.DBName + "?sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	log.Println("banco de dados conectado")
	ctx := context.Background()
	pb, err := db.Conn(ctx)
	if err != nil {
		log.Panic(err)
	}
	defer pb.Close()


}
