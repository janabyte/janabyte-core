package janabyte

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/aidosgal/janabyte/janabyte-core/internal/http/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type APIServer struct {
    address string
    db      *sql.DB
}

func NewApiServer(address string, db *sql.DB) *APIServer {
    return &APIServer{
        address: address,
        db: db,
    }
}

func (s *APIServer) Run() error {
    router := chi.NewRouter()
    router.Use(middleware.Logger)

    userHandler := handler.NewUserHandler();

    router.Route("/api/v1", func(router chi.Router) {
        router.Post("/login", userHandler.HandleLogin) 
    })

    log.Println("Listening on", s.address)

    return http.ListenAndServe(s.address, router)
}
