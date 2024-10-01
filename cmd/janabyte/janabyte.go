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

	roleRepo := repository.NewRolesRepository(s.db)
	roleService := service.NewRoleService(*roleRepo)
	roleHandler := handler.NewHandlerRole(*roleService)

	instanceRepo := repository.NewInstanceRepository(s.db)
	instanceService := service.NewServiceInstance(*instanceRepo)
	instanceHandler := handler.NewInstanceHandler(*instanceService)

	userRepository := repository.NewUserRepository(s.db)
	userService := service.NewUserService(*userRepository, *roleRepo)
	userHandler := handler.NewUserHandler(*userService)

	clubRepository := repository.NewClubRepository(s.db)
	clubService := service.NewClubService(*clubRepository, *userService)
	clubHandler := handler.NewClubHandler(*clubService)

	computerRepo := repository.NewComputerRepository(s.db)
	computerService := service.NewComputerService(*computerRepo, *clubRepository, *instanceRepo)
	computerHandler := handler.NewComputerHandler(*computerService)

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
		router.Route("/clubs", func(router chi.Router) {
			router.Get("/", clubHandler.HandlerGetClubList)
			router.Get("/{id}", clubHandler.HandlerGetClubById)
			router.Post("/", clubHandler.HandlerCreateClub)
			router.Delete("/{id}", clubHandler.HandlerDeleteClub)
			router.Put("/{id}", clubHandler.HandlerUpdateClub)
		})
		router.Route("/instances", func(router chi.Router) {
			router.Get("/", instanceHandler.HandlerGetAllInstance)
			router.Post("/", instanceHandler.HandlerCreateInstance)
			router.Get("/{id}", instanceHandler.HandlerGetInstanceById)
			router.Delete("/{id}", instanceHandler.HandlerDeleteInstanceById)
			router.Put("/{id}", instanceHandler.HandlerUpdateInstanceById)
		})
		router.Route("/roles", func(router chi.Router) {
			router.Get("/", roleHandler.HandlerGetAllRole)
			router.Post("/", roleHandler.HandlerCreateRole)
			router.Get("/{id}", roleHandler.HandlerGetRoleById)
			router.Delete("/{id}", roleHandler.HandlerDeleteRole)
			router.Put("/{id}", roleHandler.HandlerUpdateRoleById)
		})
		router.Route("/computers", func(router chi.Router) {
			router.Get("/", computerHandler.HandlerGetAllComputers)
			router.Post("/", computerHandler.HandlerCreateComputer)
			router.Get("/{id}", computerHandler.HandlerGetComputerById)
			router.Delete("/{id}", computerHandler.HandlerDeleteComputerById)
			router.Put("/{id}", computerHandler.HandlerUpdateComputerById)
		})
	})

	sloger.Info("Listening on", s.address)
	return http.ListenAndServe(s.address, router)
}
