package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	// grpcClient "github.com/renatospaka/payment-transaction/adapter/grpc/client"
	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
	httpServer "github.com/renatospaka/payment-transaction/adapter/httpServer"
	repository "github.com/renatospaka/payment-transaction/adapter/postgres"
	"github.com/renatospaka/payment-transaction/adapter/web/controller"
	"github.com/renatospaka/payment-transaction/core/usecase"
	"github.com/renatospaka/payment-transaction/utils/configs"
	"google.golang.org/grpc"
)

func main() {
	log.Println("iniciando a aplicação")
	configs, err := configs.LoadConfig(".")
	if err != nil {
		log.Panic(err)
	}

	ctx := context.Background()

	//open connection to the database
	log.Println("iniciando conexão com o banco de dados")
	conn := "postgresql://" + configs.DBUser + ":" + configs.DBPassword + "@" + configs.DBHost + ":" + configs.DBPort + "/" + configs.DBName + "?sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	log.Println("iniciando gerador de transações")
	repo := repository.NewPostgresDatabase(db)
	usecases := usecase.NewTransactionUsecase(repo)
	controllers := controller.NewTransactionController(usecases)
	webServer := httpServer.NewHttpServer(ctx, controllers)

	// client := grpcClient.NewGrpcClient(ctx)
	log.Println("estabelecendo conexão com o servidor gRPC")
	options := make([]grpc.DialOption, 0)
	client, err := grpc.Dial(":" + configs.GRPCServerPort, options...)
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()

	srv := pb.NewAuthorizationServiceClient(client)

	//start web server
	log.Printf("gerador de transações escutando porta: %s\n", configs.WEBServerPort)
	http.ListenAndServe(":"+configs.WEBServerPort, webServer.Server)
}
