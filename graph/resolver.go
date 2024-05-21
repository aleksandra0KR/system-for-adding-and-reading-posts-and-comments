package graph

import "system-for-adding-and-reading-posts-and-comments/innternal/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Repository *repository.Repository
}
