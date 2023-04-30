package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/0xYami/twitter/models"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type profileRouter struct{}

func (rs profileRouter) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.With(profileContext).Get("/", rs.Get)
	})

	return r
}

type profileResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func profileContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("twitter-token")
		if err != nil {
			http.Error(w, "cookie not found", http.StatusUnauthorized)
			return
		}

		db := r.Context().Value("db").(*gorm.DB)
		user := &models.User{}

		if err := db.Where("token = ?", cookie.Value).First(&user).Error; err != nil {
			http.Error(w, "failed to get user", http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (rs profileRouter) Get(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*models.User)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
