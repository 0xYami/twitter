package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/0xYami/twitter/models"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type tweetRouter struct{}

func (rs tweetRouter) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.With(profileContext).Post("/", rs.Create)

		r.Get("/latest", rs.ListLatest)
	})

	return r
}

type tweetUser struct {
	Name   string `json:"name"`
	Handle string `json:"handle"`
}

type tweetResponse struct {
	ID        uint      `json:"id"`
	Text      string    `json:"text"`
	User      tweetUser `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
}

func (rs tweetRouter) ListLatest(w http.ResponseWriter, r *http.Request) {
	var tweets []models.Tweet
	db := r.Context().Value("db").(*gorm.DB)
	if err := db.Preload("User").Model(&models.Tweet{}).Find(&tweets).Error; err != nil {
		http.Error(w, "Failed to retrieve latest tweets", http.StatusInternalServerError)
		return
	}

	res := make([]tweetResponse, len(tweets))

	for i, tweet := range tweets {
		res[i] = tweetResponse{
			ID:        tweet.ID,
			Text:      tweet.Text,
			CreatedAt: tweet.CreatedAt,
			User: tweetUser{
				Name:   tweet.User.Username,
				Handle: tweet.User.Handle,
			},
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

type createTweetRequest struct {
	Text string `json:"text" validate:"required,min=0,max=280"`
}

type createTweetResponse struct {
	ID   uint   `json:"id"`
	Text string `json:"text"`
}

func (rs tweetRouter) Create(w http.ResponseWriter, r *http.Request) {
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

	res := &createTweetResponse{
		ID:   tweet.UserID,
		Text: tweet.Text,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
