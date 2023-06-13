package server

import (
	"context"
	"log"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type HttpServer struct {
	ctx    context.Context
	Server *chi.Mux
}

func NewHttpServer(ctx context.Context) *HttpServer {
	log.Println("iniciando servidor http")
	httpServer := &HttpServer{
		ctx: ctx,
	}
	httpServer.Server = connect()

	return httpServer
}

func connect() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	// router.Use(middleware.WithValue("jwt", configs.TokenAuth))
	// router.Use(middleware.WithValue("JWTExpiresIn", configs.JWTExpiresIn))

	router.Route("/transactions", func(r chi.Router) {
		// // r.Use(jwtauth.Verifier(configs.TokenAuth))
		// // r.Use(jwtauth.Authenticator)

		// r.Post("/", ProductHandler.ProcessTransaction)
		// r.Get("/", ProductHandler.GetAllTransactions)
		// r.Get("/{id}", ProductHandler.GetTransaction)
		// r.Put("/{id}", ProductHandler.ModifyTransaction)
		// r.Delete("/{id}", ProductHandler.RemoveTransaction)
	})

	return router
}
