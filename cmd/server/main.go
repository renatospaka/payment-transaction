package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/renatospaka/transact/configs"
)

func main(){
	log.Println("iniciando a aplicação")
	configs, err := configs.LoadConfig(".")
	if err != nil {
		log.Panic(err)
	}

	// log.Printf("DBHost: %s\n", configs.DBHost)
	// log.Printf("DBUser: %s\n", configs.DBUser)
	// log.Printf("DBPassword: %s\n", configs.DBPassword)
	// log.Printf("DBPort: %s\n", configs.DBPort)
	// log.Printf("WEBServerPort: %s\n", configs.WEBServerPort)
	// log.Printf("GRPCServerPort: %s\n", configs.GRPCServerPort)
	
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
