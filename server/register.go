package server

import (
	"encoding/json"
	"net/http"

	"github.com/0xYami/twitter/models"
	"gorm.io/gorm"
)

type registerRequest struct {
	Username string `json:"username" validate:"required,min=3,max=15"`
	Password string `json:"password" validate:"required,min=8"`
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	var rr registerRequest

	if err := json.NewDecoder(r.Body).Decode(&rr); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := models.NewUser(rr.Username, rr.Password)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	db := r.Context().Value("db").(*gorm.DB)
	if err := db.Create(&user).Error; err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	res := &profileResponse{
		ID:       user.ID,
		Username: user.Username,
		Token:    user.Token,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
