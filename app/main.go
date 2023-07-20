package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"

	"github.com/renatospaka/payment-transaction/adapter/grpc/client"
	httpServer "github.com/renatospaka/payment-transaction/adapter/httpServer"
	repository "github.com/renatospaka/payment-transaction/adapter/postgres"
	"github.com/renatospaka/payment-transaction/adapter/web/controller"
	"github.com/renatospaka/payment-transaction/core/service"
	"github.com/renatospaka/payment-transaction/core/usecase"
	"github.com/renatospaka/payment-transaction/utils/configs"
)

func main() {
	log.Println("iniciando a aplicação")
	configs, err := configs.LoadConfig(".")
	if err != nil {
		log.Panic(err)
	}

	deadLineWEB := time.Now().Add(time.Duration(configs.WEBServerTimeOut) * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadLineWEB)
	defer cancel()

	//open connection to the database
	log.Println("iniciando conexão com o banco de dados")
	conn := "postgresql://" + configs.DBUser + ":" + configs.DBPassword + "@" + configs.DBHost + ":" + configs.DBPort + "/" + configs.DBName + "?sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	// grpc server
	log.Printf("iniciando conexão com o servidor gRPC na porta :%s\n", configs.GRPCServerPort)
	options := make([]grpc.DialOption, 0)
	options = append(options, grpc.WithInsecure())
	connGrpc, err := grpc.Dial(":"+configs.GRPCServerPort, options...)
	if err != nil {
		log.Panic(err)
	}
	defer connGrpc.Close()

	log.Println("iniciando gerador de transações")

	//web server
	repo := repository.NewPostgresDatabase(db)
	usecases := usecase.NewTransactionUsecase(repo)
	controllers := controller.NewTransactionController(usecases)
	webServer := httpServer.NewHttpServer(ctx, controllers)

	//grpc services
	clientGrpc := client.NewGrpcClient(ctx, connGrpc)
	services := service.NewTransactionService(clientGrpc)
	usecases.SetServices(services)

	//start web server
	log.Printf("gerador de transações escutando porta: %s\n", configs.WEBServerPort)
	http.ListenAndServe(":"+configs.WEBServerPort, webServer.Server)
}
