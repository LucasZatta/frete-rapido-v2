package repository

import (
	"github.com/lucaszatta/frete-rapido-v2/internal/quote/models"
	"gorm.io/gorm"
)

type QuoteRepository interface {
	Create(quote *models.Quote) error
	GetByID(id int) (*models.Quote, error)
}

type quoteRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *quoteRepository {
	return &quoteRepository{
		db: db,
	}
}

func (p *quoteRepository) Create(quote *models.Quote) error {
	tx := p.db.Create(quote)
	return tx.Error
}

func (p *quoteRepository) GetByID(id int) (*models.Quote, error) {
	quote := &models.Quote{}

	tx := p.db.Where("id = ?", id).First(quote)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return quote, nil
}
