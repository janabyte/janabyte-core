package janabyte

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/aidosgal/janabyte/janabyte-core/internal/http/handler"
	"github.com/gorilla/mux"
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
    router := mux.NewRouter()
    subrouter := router.PathPrefix("/api/v1").Subrouter()

    userHandler := handler.NewUserHandler()
    userHandler.RegisterRoutes(subrouter)

    log.Println("Listening on", s.address)

    return http.ListenAndServe(s.address, router)
}
