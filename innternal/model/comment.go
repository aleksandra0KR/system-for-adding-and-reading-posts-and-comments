package model

import (
	"github.com/google/uuid"
	"time"
)

type Comment struct {
	CommentId uuid.UUID  `json:"commentId,omitempty"`
	Body      string     `json:"body,omitempty"`
	UserId    uuid.UUID  `json:"userId,omitempty"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt"`
	Parent    *Comment   `json:"parent,omitempty"`
	Children  []*Comment `json:"children"`
	Post      *Post      `json:"post,omitempty"`
}
