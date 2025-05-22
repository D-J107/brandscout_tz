package repositories

import "quotes_api/internal/app/domain/models"

type QuotesRepository interface {
	Save(author, content string) (*models.Quote, error)
	GetAll() ([]*models.Quote, error)
	GetById(id int) (*models.Quote, error)
	GetByAuthor(author string) ([]*models.Quote, error)
	Delete(id int) error
}

func NewQuotesRepository() QuotesRepository {
	return &quotesRepositorySimpleImplementation{storage: make([]*models.Quote, 0)}
}
