package httpServer

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc/codes"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/renatospaka/payment-transaction/adapter/web/controller"
	middlewares "github.com/renatospaka/payment-transaction/adapter/web/middleware"
)

type HttpServer struct {
	ctx         context.Context
	controllers *controller.TransactionController
	Server      *chi.Mux
}

func NewHttpServer(ctx context.Context, controller *controller.TransactionController) *HttpServer {
	log.Println("iniciando conexão com o servidor web")
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
	s.Server.Use(middlewares.Cors)
	// s.Server.Use(middleware.WithValue("jwt", configs.TokenAuth))
	// s.Server.Use(middleware.WithValue("JWTExpiresIn", configs.JWTExpiresIn))

	s.Server.Route("/transactions", func(r chi.Router) {
		// // r.Use(jwtauth.Verifier(configs.TokenAuth))
		// // r.Use(jwtauth.Authenticator)

		time.Sleep(400 * time.Millisecond)
		if 	s.ctx.Err() == context.Canceled {
			log.Fatalf("error processing the RPC call: %v\n", codes.Canceled)
		}

		r.Post("/", s.controllers.Process)
		r.Post("/{transactioId}", s.controllers.ReprocessPending)
		r.Get("/{id}", s.controllers.Get)
		r.Get("/", s.controllers.GetAll)
		r.Put("/{id}", s.controllers.Modify)
		r.Delete("/{id}", s.controllers.Remove)
	})

	s.Server.Route("/health", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode("Healthy")
		})
	})
}
