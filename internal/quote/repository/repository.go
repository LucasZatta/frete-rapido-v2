package repository

import (
	"fmt"

	"github.com/lucaszatta/frete-rapido-v2/internal/quote/models"
	"gorm.io/gorm"
)

type QuoteRepository interface {
	CreateSingle(quote *models.Quote) error
	GetByID(id int) (*models.Quote, error)
	Create(quotes *[]models.Quote) error
	GetLastQuotes(lastQuotes string) (*[]models.LastQuotes, error)
	GetMaxMinQuotes() (*models.Quote, *models.Quote, error)
}

type quoteRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *quoteRepository {
	return &quoteRepository{
		db: db,
	}
}

func (p *quoteRepository) CreateSingle(quote *models.Quote) error {
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

func (p *quoteRepository) Create(quotes *[]models.Quote) error {
	tx := p.db.Create(quotes)
	return tx.Error
}

func (p *quoteRepository) GetLastQuotes(lastQuotes string) (*[]models.LastQuotes, error) {
	quotes := make([]models.LastQuotes, 0)
	var query string
	if lastQuotes == "" {
		query = "SELECT name, count(*) count, sum(price) price_sum, ROUND(sum(price)/count(*), 2) price_avg FROM (SELECT * FROM public.quotes ) GROUP BY name "
	} else {
		query = fmt.Sprintf("SELECT name, count(*) count, sum(price) price_sum, ROUND(sum(price)/count(*), 2) price_avg FROM (SELECT * FROM public.quotes LIMIT %s ) GROUP BY name", lastQuotes)
	}

	tx := p.db.Raw(query).Scan(&quotes)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &quotes, nil
}

func (p *quoteRepository) GetMaxMinQuotes() (*models.Quote, *models.Quote, error) {
	maxQuote, minQuote := &models.Quote{}, &models.Quote{}

	tx := p.db.Order("price asc").First(&minQuote)
	if tx.Error != nil {
		return nil, nil, tx.Error
	}
	tx = p.db.Order("price desc").First(&maxQuote)
	if tx.Error != nil {
		return nil, nil, tx.Error
	}

	return maxQuote, minQuote, nil
}
