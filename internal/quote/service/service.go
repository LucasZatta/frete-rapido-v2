package service

import (
	"github.com/lucaszatta/frete-rapido-v2/internal/quote/models"
	"github.com/lucaszatta/frete-rapido-v2/internal/quote/repository"
)

type QuoteService interface {
	Create(quote *[]models.Quote) (*[]models.Quote, error)
	GetLastQuotes(lastQuotes string) (*[]models.LastQuotes, error)
	GetMaxMinQuotes() (*models.Quote, *models.Quote, error)
}

type quoteService struct {
	quoteRepository repository.QuoteRepository
}

func New(quoteRepository repository.QuoteRepository) *quoteService {
	return &quoteService{
		quoteRepository: quoteRepository,
	}
}

func (p *quoteService) Create(quotes *[]models.Quote) (*[]models.Quote, error) {
	err := p.quoteRepository.Create(quotes)

	if err != nil {
		return nil, err
	}
	return quotes, nil
}

func (p *quoteService) GetLastQuotes(lastQuotes string) (*[]models.LastQuotes, error) {
	return p.quoteRepository.GetLastQuotes(lastQuotes)
}

func (p *quoteService) GetMaxMinQuotes() (*models.Quote, *models.Quote, error) {
	return p.quoteRepository.GetMaxMinQuotes()
}
