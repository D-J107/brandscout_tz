package repositories

import (
	"quotes_api/internal/app/domain/models"
)

type quotesRepositorySimpleImplementation struct {
	storage []*models.Quote
}

func (r *quotesRepositorySimpleImplementation) Save(author, content string) (*models.Quote, error) {
	newQuote := &models.Quote{}
	if len(r.storage) == 0 {
		newQuote = &models.Quote{Id: 1, Author: author, Content: content}
	} else {
		newQuote = &models.Quote{Id: r.storage[len(r.storage)-1].Id + 1, Author: author, Content: content}
	}
	r.storage = append(r.storage, newQuote)
	return newQuote, nil
}

func (r *quotesRepositorySimpleImplementation) GetAll() ([]*models.Quote, error) {
	return r.storage, nil
}

func (r *quotesRepositorySimpleImplementation) GetById(id int) (*models.Quote, error) {
	if id <= 0 {
		return nil, ErrInvalidId
	}
	for _, quote := range r.storage {
		if quote.Id == id {
			return quote, nil
		}
	}
	return nil, nil
}

func (r *quotesRepositorySimpleImplementation) GetByAuthor(author string) ([]*models.Quote, error) {
	if author == "" {
		return nil, ErrEmptyAuthor
	}
	authorQuotes := make([]*models.Quote, 0)
	for _, quote := range r.storage {
		if quote.Author == author {
			authorQuotes = append(authorQuotes, quote)
		}
	}
	return authorQuotes, nil
}

func (r *quotesRepositorySimpleImplementation) Delete(id int) error {
	if id <= 0 {
		return ErrInvalidId
	}
	for i := range r.storage {
		if r.storage[i].Id == id {
			r.storage = remove(r.storage, i)
			return nil
		}
	}
	return nil
}

func remove(s []*models.Quote, i int) []*models.Quote {
	return append(s[:i], s[i+1:]...)
}
