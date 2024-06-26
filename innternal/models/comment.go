package models

import (
	"github.com/google/uuid"
)

type Comment struct {
	Id       uuid.UUID  `json:"id,omitempty"`
	Body     string     `json:"body,omitempty"`
	UserId   uuid.UUID  `json:"userId,omitempty"`
	Parent   uuid.UUID  `json:"parent,omitempty"`
	Children []*Comment `json:"children"`
	Post     uuid.UUID  `json:"postId,omitempty"`
}
