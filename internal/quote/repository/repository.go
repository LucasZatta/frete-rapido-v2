package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type QuoteRepository interface {
	GetQuote()
}

type quoteRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *quoteRepository {
	return &quoteRepository{
		db: db,
	}
}

func (p *quoteRepository) GetQuote() {
	fmt.Println("hello")
}
