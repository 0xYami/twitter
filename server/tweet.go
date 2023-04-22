package server

import (
	"encoding/json"
	"net/http"

	"github.com/0xYami/twitter/models"
	"gorm.io/gorm"
)

func getUserTweets(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*models.User)
	db := r.Context().Value("db").(*gorm.DB)

	var tweets []models.Tweet
	if err := db.Where("user_id = ?", user.ID).Find(&tweets).Error; err != nil {
		http.Error(w, "Failed to get tweets", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweets)
}

type createTweetRequest struct {
	Text string `json:"text" validate:"required,min=3,max=140"`
}

func createTweet(w http.ResponseWriter, r *http.Request) {
	nt := &createTweetRequest{}

	if err := json.NewDecoder(r.Body).Decode(&nt); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user := r.Context().Value("user").(*models.User)
	tweet := &models.Tweet{
		Text:   nt.Text,
		UserID: user.ID,
		User:   *user,
	}

	db := r.Context().Value("db").(*gorm.DB)
	if err := db.Create(&tweet).Error; err != nil {
		http.Error(w, "Failed to create tweet", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tweet)
}
