package server

import (
	"encoding/json"
	"net/http"

	"github.com/0xYami/twitter/models"
	"gorm.io/gorm"
)

func auth(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("twitter-token")
	if err != nil {
		http.Error(w, "cookie not found", http.StatusUnauthorized)
		return
	}

	db := r.Context().Value("db").(*gorm.DB)
	user := &models.User{}

	if err := db.Where("token = ?", cookie.Value).First(&user).Error; err != nil {
		http.Error(w, "failed to get user", http.StatusUnauthorized)
		return
	}

	res := &profileResponse{
		ID:        user.ID,
		Username:  user.Username,
		Handle:    user.Handle,
		AvatarURL: user.AvatarURL,
		Token:     user.Token,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
