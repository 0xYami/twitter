package server

import (
	"context"
	"net/http"

	"github.com/0xYami/twitter/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(middleware.CleanPath)
	s.Router.Use(middleware.AllowContentType("application/json"))
	s.Router.Use(middleware.Heartbeat("/_health"))
	s.Router.Use(middleware.Logger)
	s.Router.Use(s.DBMiddleware)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.Timeout(s.Config.Timeout))

	s.Router.Mount("/debug", middleware.Profiler())

	s.Router.Route("/users", func(r chi.Router) {
		r.Get("/", getAllUsers)
		r.Post("/", createUser)

		r.Route("/{userID}", func(r chi.Router) {
			r.Use(userContext)
			r.Get("/", getUserById)
			r.Get("/tweets", getUserTweets)
			r.Post("/tweets", createTweet)
		})
	})
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.Config.Address, s.Router)
}
