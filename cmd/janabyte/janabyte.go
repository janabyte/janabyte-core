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
			//router.Post("/register", handler.CreateUser(userRepository))
			//router.Get("/all", handler.GetAllUsers(userRepository))
			//router.Get("/{id}", handler.GetUserById(userRepository))
			//router.Delete("/delete/{id}", handler.DeleteUserById(userRepository))
			//router.Put("/update", handler.UpdateUserById(userRepository))
		})
	})

	//log.Println("Listening on", s.address)
	sloger.Info("Listening on", s.address)
	return http.ListenAndServe(s.address, router)
}
