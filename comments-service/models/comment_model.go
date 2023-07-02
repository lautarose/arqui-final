package models

import "time"

type Comment struct {
	CommentID int       `gorm:"primary_key"`
	UserID    string    `gorm:"type:varchar(255);not null"`
	ItemID    string    `gorm:"type:varchar(255);not null"`
	Message   string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"not null"`
}

type Comments []Comment
