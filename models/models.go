package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"unique;not null"`
	Handle    string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Token     string `gorm:"unique;not null"`
	Bio       string
	AvatarURL string
	Tweets    []Tweet    `gorm:"foreignKey:UserID"`
	Followers []Follower `gorm:"foreignKey:FollowerID"`
	Following []Follower `gorm:"foreignKey:FollowingID"`
}

func NewUser(username string, password string) (*User, error) {
	user := &User{
		Username: username,
		Password: password,
		Handle:   username,
	}

	token, err := user.generateJWT()
	if err != nil {
		return nil, err
	}

	user.Token = token
	return user, nil
}

func (u *User) generateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":   u.Username,
		"password":   u.Password,
		"expiration": time.Now().Add(24 * time.Hour),
	})
	return token.SignedString([]byte("secret"))
}

type Tweet struct {
	gorm.Model
	Text   string `gorm:"not null"`
	UserID uint   `gorm:"not null"`
	User   User   `gorm:"foreignKey:UserID"`
}

type Follower struct {
	ID          uint `gorm:"primaryKey"`
	FollowerID  uint `gorm:"not null"`
	FollowingID uint `gorm:"not null"`
	Follower    User `gorm:"foreignKey:FollowerID"`
	Following   User `gorm:"foreignKey:FollowingID"`
}
