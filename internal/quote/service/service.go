package service

import (
	"github.com/lucaszatta/frete-rapido-v2/internal/quote/repository"
)

type QuoteService interface {
	GetQuote()
}

type quoteService struct {
	quoteRepository repository.QuoteRepository
}

func New(quoteRepository repository.QuoteRepository) *quoteService {
	return &quoteService{
		quoteRepository: quoteRepository,
	}
}

func (p *quoteService) GetQuote() {
	p.quoteRepository.GetQuote()
}
