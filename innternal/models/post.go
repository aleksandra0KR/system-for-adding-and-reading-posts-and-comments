package models

import (
	"github.com/google/uuid"
)

type Post struct {
	Id       uuid.UUID  `json:"id,omitempty"`
	Title    string     `json:"title,omitempty"`
	Body     string     `json:"body,omitempty"`
	UserId   uuid.UUID  `json:"user_id,omitempty"`
	Comments []*Comment `json:"comments,omitempty"`
	Disabled bool       `json:"disabled,omitempty"`
}
