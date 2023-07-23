package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	ctx := context.Background()

	//open connection to the database
	log.Println("iniciando conexão com o banco de dados")
	// conn := "postgresql://" + configs.DBUser + ":" + configs.DBPassword + "@" + configs.DBHost + ":" + string(configs.DBPort) + "/" + configs.DBName + "?sslmode=disable"
	conn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	// grpc server
	log.Printf("iniciando conexão com o servidor gRPC na porta :%d\n", configs.GRPCServerPort)
	options := make([]grpc.DialOption, 0)
	options = append(options, grpc.WithInsecure())
	connGrpc, err := grpc.Dial(":" + string(configs.GRPCServerPort), options...)
	if err != nil {
		log.Panic(err)
	}
	defer connGrpc.Close()

	log.Println("iniciando gerador de transações")

	// Prepare the Domain to start up
	repo := repository.NewPostgresDatabase(db)
	usecases := usecase.NewTransactionUsecase(repo)
	controllers := controller.NewTransactionController(usecases)
	webServer := httpServer.NewHttpServer(ctx, controllers)

	// Prepare the HTTP Server to connect
	serverAddress := fmt.Sprintf("%s:%d", configs.WEBServerHost, configs.WEBServerPort)
	server := &http.Server{
		Addr: serverAddress,
		Handler:  webServer.Server,
	}

	// Prepare the gRPC client to connect to the gRPC Server
	clientGrpc := client.NewGrpcClient(ctx, connGrpc)
	services := service.NewTransactionService(clientGrpc)
	usecases.SetServices(services)

	// Prepare graceful shutdown
	serverCtx, serverStop := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func(){
		<- sign

		shutdownCtx, _ := context.WithTimeout(serverCtx, 500 * time.Millisecond)
		go func(){
			<- shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("tempo normal de desligamento da aplicação esgotado... forçando o desligamento!")
			}
			
			// Trigger graceful shutdown
			err = server.Shutdown(shutdownCtx)
			if err != nil {
				log.Fatalf("desligamento normal da aplicação com erro: %v\n", err)
				log.Fatal("forçando o desligamento!")
				serverStop()
			}
		}()
	}()

	// Start up WEB Server
	log.Printf("gerador de transações escutando em: %s\n", serverAddress)
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("erro quando a aplicação começou a escutar e servir: %v\n", err)
	}

	// Wait fot server context to be stopped
	<- serverCtx.Done()
}
