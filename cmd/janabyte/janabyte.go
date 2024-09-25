package janabyte

import (
	"database/sql"
	"github.com/aidosgal/janabyte/janabyte-core/internal/logger"
	"github.com/aidosgal/janabyte/janabyte-core/internal/service"
	"net/http"

	"github.com/aidosgal/janabyte/janabyte-core/internal/http/handler"
	"github.com/aidosgal/janabyte/janabyte-core/internal/repository"
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
		db:      db,
	}
}

func (s *APIServer) Run() error {
	sloger := logger.SetupLogger()
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.URLFormat)
	userRepository := repository.NewUserRepository(s.db)
	userService := service.NewUserService(*userRepository)
	userHandler := handler.NewUserHandler(*userService)

	router.Route("/api/v1", func(router chi.Router) {
		router.Route("/users", func(router chi.Router) {
			router.Get("/", userHandler.HandleGetAllUsers)
			router.Post("/", userHandler.HandleCreateUser)
			router.Delete("/{id}", userHandler.HandleDeleteUserById)
			router.Put("/{id}", userHandler.HandleUpdateUserById)
			router.Get("/{id}", userHandler.HandleGetUserById)
			router.Post("/login", userHandler.HandleLogin)
			router.Post("/logout", userHandler.Logout)
			router.Post("/refresh", userHandler.RefreshTokenHandler)
		})
	})

	sloger.Info("Listening on", s.address)
	return http.ListenAndServe(s.address, router)
}
