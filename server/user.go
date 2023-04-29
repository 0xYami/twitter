package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/0xYami/twitter/models"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type userResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func userContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("twitter-token")
		if err != nil {
			http.Error(w, "[context] cookie not found", http.StatusUnauthorized)
			return
		}

		id := chi.URLParam(r, "id")
		if id == "" {
			http.Error(w, "[context] missing user id", http.StatusUnauthorized)
			return
		}

		db := r.Context().Value("db").(*gorm.DB)
		user := &models.User{}

		if err := db.Where("id = ? AND token = ?", id, cookie.Value).First(&user).Error; err != nil {
			http.Error(w, "[context] failed to get user", http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*models.User)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

type createUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=15"`
	Password string `json:"password" validate:"required,min=8"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var nu createUserRequest

	if err := json.NewDecoder(r.Body).Decode(&nu); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := models.NewUser(nu.Username, nu.Password)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	db := r.Context().Value("db").(*gorm.DB)
	if err := db.Create(&user).Error; err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	res := &userResponse{
		ID:       user.ID,
		Username: user.Username,
		Token:    user.Token,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
