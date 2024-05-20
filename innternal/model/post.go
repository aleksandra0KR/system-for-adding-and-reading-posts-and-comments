package model

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	PostId           uuid.UUID  `json:"postId,omitempty"`
	Title            string     `json:"title,omitempty"`
	Body             string     `json:"body,omitempty"`
	UserId           uuid.UUID  `json:"userId,omitempty"`
	Comments         []*Comment `json:"comments,omitempty"`
	DisabledComments bool       `json:"disabledComments,omitempty"`
	CreatedAt        time.Time  `json:"createdAt,omitempty"`
	UpdatedAt        time.Time  `json:"updatedAt"`
}
