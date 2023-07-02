package dto

import "time"

type CommentDto struct {
	CommentID int       `json:"comment_id"`
	UserID    string    `json:"user_id"`
	ItemID    string    `json:"item_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentsDto []CommentDto
