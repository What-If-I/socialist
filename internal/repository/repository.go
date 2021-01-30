package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	User userRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: userRepo{db: db},
	}
}
