package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"unique;not null"`
	Handle    string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Bio       string
	AvatarURL string
	Tweets    []Tweet    `gorm:"foreignKey:UserID"`
	Followers []Follower `gorm:"foreignKey:FollowerID"`
	Following []Follower `gorm:"foreignKey:FollowingID"`
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
