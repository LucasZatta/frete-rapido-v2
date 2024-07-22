package service

import (
	"github.com/lucaszatta/frete-rapido-v2/internal/quote/models"
	"github.com/lucaszatta/frete-rapido-v2/internal/quote/repository"
)

type QuoteService interface {
	Create(quote *models.Quote) (*models.Quote, error)
	GetByID(id int) (*models.Quote, error)
}

type quoteService struct {
	quoteRepository repository.QuoteRepository
}

func New(quoteRepository repository.QuoteRepository) *quoteService {
	return &quoteService{
		quoteRepository: quoteRepository,
	}
}

func (p *quoteService) Create(quote *models.Quote) (*models.Quote, error) {

	err := p.quoteRepository.Create(quote)
	if err != nil {
		return nil, err
	}
	return quote, nil
}

func (p *quoteService) GetByID(id int) (*models.Quote, error) {

	return p.quoteRepository.GetByID(id)
}
