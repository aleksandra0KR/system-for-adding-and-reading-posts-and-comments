package graph

import (
	"system-for-adding-and-reading-posts-and-comments/innternal/repository"
)

type Resolver struct {
	Repository repository.Database
}
