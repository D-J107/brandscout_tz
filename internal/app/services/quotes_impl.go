package services

import (
	"math/rand"
	"quotes_api/internal/app/domain/models"
	"quotes_api/internal/app/repositories"
	"quotes_api/internal/logging"
)

type QuotesServiceImplementation struct {
	repo repositories.QuotesRepository
}

func (qs *QuotesServiceImplementation) Save(author, content string) (*models.Quote, error) {
	createdQuote, err := qs.repo.Save(author, content)
	if err != nil {
		logging.Logger.Error("QuotesService.Save() ", "error", err)
		return nil, err
	}
	return createdQuote, nil
}

func (qs *QuotesServiceImplementation) GetAll() ([]*models.Quote, error) {
	quotes, err := qs.repo.GetAll()
	if err != nil {
		logging.Logger.Error("QuotesService.GetAll() ", "error", err)
		return nil, err
	}
	return quotes, nil
}

func (qs *QuotesServiceImplementation) GetRandom() (*models.Quote, error) {
	quotes, err := qs.repo.GetAll()
	if err != nil {
		logging.Logger.Error("QuotesService.GetRandom() ", "error", err)
		return nil, err
	}
	return quotes[rand.Intn(len(quotes))], nil
}

func (qs *QuotesServiceImplementation) GetByAuthor(author string) ([]*models.Quote, error) {
	quotes, err := qs.repo.GetByAuthor(author)
	if err != nil {
		logging.Logger.Error("QuotesService.GetByAuthor() ", "error", err)
		return nil, err
	}
	return quotes, nil
}

func (qs *QuotesServiceImplementation) DeleteById(id int) error {
	err := qs.repo.Delete(id)
	if err != nil {
		logging.Logger.Error("QuotesService.DeleteById() ", "error", err)
	}
	return err
}
