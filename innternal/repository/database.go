package repository

import (
	"github.com/google/uuid"
	"system-for-adding-and-reading-posts-and-comments/graph/model"
	"system-for-adding-and-reading-posts-and-comments/innternal/models"
)

const (
	host     = "ozon-contest"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "OzonContest"
)

type Database interface {
	CreateComment(comment *models.Comment) (*models.Comment, error)
	DeleteCommentByID(id uuid.UUID) error
	GetCommentsForPost(id uuid.UUID, limit, offset int) ([]*model.Comment, error)
	UpdateComment(comment *models.Comment) (*models.Comment, error)
	CreatePost(post *models.Post) (*models.Post, error)
	DeletePostByID(id uuid.UUID) error
	GetPostByID(id uuid.UUID) (*models.Post, error)
	UpdatePost(post *models.Post) (*models.Post, error)
	CreateUser(user *models.User) (*models.User, error)
	DeleteUserByID(id uuid.UUID) error
}

/*
func NewDataBase(databaseModel string) graph.Config {
	if databaseModel == "in-memory" {
		repo := inMemory.NewInMemoryRepository()
		return graph.Config{Resolvers: &graph.Resolver{Repository: repo}}
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	fmt.Println("Connectiong")
	if err != nil {
		log.Fatalf("connection failed 1: %s", err.Error())
	}
	fmt.Println("Connectiong2")
	defer db.Close()
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	migration, err := migrate.NewWithDatabaseInstance("file://database/migration",
		dbname, driver)
	if err != nil {
		log.Fatal(err)
	}
	if err := migration.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal(err)
		}
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	repo := postgres2.NewRepository(db)
	return graph.Config{Resolvers: &graph.Resolver{Repository: repo}}

}
*/
