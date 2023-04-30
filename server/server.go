package server

import (
	"context"
	"net/http"

	"github.com/0xYami/twitter/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	Router *chi.Mux
	DB     *gorm.DB
	Logger *zap.Logger
	Config *config.Config
}

func CreateNewServer(cfg *config.Config, logger *zap.Logger) *Server {
	return &Server{
		Router: chi.NewRouter(),
		Logger: logger,
		Config: cfg,
	}
}

func (s *Server) InitDB() error {
	dsn := s.Config.DB.DSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		QueryFields: true,
	})
	if err != nil {
		return err
	}
	s.DB = db
	return nil
}

func (s *Server) DBMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), s.Config.Timeout)
		defer cancel()

		db := s.DB.WithContext(ctx)
		ctx = context.WithValue(r.Context(), "db", db)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func (s *Server) MountHandlers() {
	s.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}))
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(middleware.CleanPath)
	s.Router.Use(middleware.AllowContentType("application/json"))
	s.Router.Use(middleware.Heartbeat("/_health"))
	s.Router.Use(middleware.Logger)
	s.Router.Use(s.DBMiddleware)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.Timeout(s.Config.Timeout))

	apiRouter := chi.NewRouter()

	apiRouter.Mount("/debug", middleware.Profiler())

	apiRouter.Post("/auth", auth)
	apiRouter.Post("/register", registerUser)

	apiRouter.Mount("/profiles/{id}", profileRouter{}.Routes())
	apiRouter.Mount("/tweets", tweetRouter{}.Routes())

	s.Router.Mount("/api", apiRouter)
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.Config.Address, s.Router)
}
