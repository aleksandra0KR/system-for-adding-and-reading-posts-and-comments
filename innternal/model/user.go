package model

import "github.com/google/uuid"

type User struct {
	UserId   uuid.UUID  `json:"userId,omitempty"`
	Posts    []*Post    `json:"posts,omitempty"`
	Comments []*Comment `json:"comments,omitempty"`
}
