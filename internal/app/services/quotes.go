package services

import (
	"quotes_api/internal/app/domain/models"
	"quotes_api/internal/app/repositories"
)

type QuotesService interface {
	Save(author, content string) (*models.Quote, error)
	GetAll() ([]*models.Quote, error)
	GetRandom() (*models.Quote, error)
	GetByAuthor(author string) ([]*models.Quote, error)
	DeleteById(id int) error
}

func NewQuoteService() QuotesService {
	return &QuotesServiceImplementation{repo: repositories.NewQuotesRepository()}
}
