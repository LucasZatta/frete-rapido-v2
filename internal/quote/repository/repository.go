package repository

import (
	"database/sql"
	"fmt"
)

type QuoteRepository interface {
	GetQuote()
}

type quoteRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *quoteRepository {
	return &quoteRepository{
		db: db,
	}
}

func (p *quoteRepository) GetQuote() {
	fmt.Println("hello")
}
