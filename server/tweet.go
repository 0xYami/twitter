package server

import (
	"encoding/json"
	"net/http"

	"github.com/0xYami/twitter/models"
	"gorm.io/gorm"
)

type createTweetRequest struct {
	Text string `json:"text" validate:"required,min=0,max=280"`
}

type tweetResponse struct {
	ID   uint   `json:"id"`
	Text string `json:"text"`
}

func createTweet(w http.ResponseWriter, r *http.Request) {
	var nt createTweetRequest

	if err := json.NewDecoder(r.Body).Decode(&nt); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user := r.Context().Value("user").(*models.User)

	tweet := &models.Tweet{
		Text:   nt.Text,
		UserID: user.ID,
	}

	db := r.Context().Value("db").(*gorm.DB)
	if err := db.Create(&tweet).Error; err != nil {
		http.Error(w, "Failed to create tweet", http.StatusInternalServerError)
		return
	}

	res := &tweetResponse{
		ID:   tweet.UserID,
		Text: tweet.Text,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
