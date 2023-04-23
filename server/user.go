package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/0xYami/twitter/models"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func userContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := r.Context().Value("db").(*gorm.DB)
		id := chi.URLParam(r, "userID")
		user := &models.User{}

		if err := db.Where("id = ?", id).First(&user).Error; err != nil {
			http.Error(w, "[context] failed to get user", http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*models.User)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db := r.Context().Value("db").(*gorm.DB)
	if err := db.Find(&users).Error; err != nil {
		http.Error(w, "Failed to get users", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
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

	user := &models.User{
		Username: nu.Username,
		Password: nu.Password,
	}

	db := r.Context().Value("db").(*gorm.DB)
	if err := db.Create(&user).Error; err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
