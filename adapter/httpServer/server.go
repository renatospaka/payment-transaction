package httpServer

import (
	"context"
	"log"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/renatospaka/payment-transaction/adapter/rest/controller"
)

type HttpServer struct {
	ctx         context.Context
	controllers *controller.TransactionController
	Server      *chi.Mux
}

func NewHttpServer(ctx context.Context, controller *controller.TransactionController) *HttpServer {
	log.Println("iniciando servidor http")
	httpServer := &HttpServer{
		ctx:         ctx,
		controllers: controller,
	}
	httpServer.connect()

	return httpServer
}

func (s *HttpServer) connect() {
	s.Server = chi.NewRouter()
	s.Server.Use(middleware.Logger)
	s.Server.Use(middleware.Recoverer)
	// s.Server.Use(middleware.WithValue("jwt", configs.TokenAuth))
	// s.Server.Use(middleware.WithValue("JWTExpiresIn", configs.JWTExpiresIn))

	s.Server.Route("/transactions", func(r chi.Router) {
		// // r.Use(jwtauth.Verifier(configs.TokenAuth))
		// // r.Use(jwtauth.Authenticator)

		r.Post("/", s.controllers.Process)
		r.Get("/", s.controllers.GetAll)
		r.Get("/{id}", s.controllers.Get)
		r.Put("/{id}", s.controllers.Modify)
		r.Delete("/{id}", s.controllers.Remove)
	})
}
