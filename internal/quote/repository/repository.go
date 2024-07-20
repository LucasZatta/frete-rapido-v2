package repository

import "gorm.io/gorm"

type QuoteRepository interface {
}

type quoteRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *quoteRepository {
	return &quoteRepository{
		db: db,
	}
}
