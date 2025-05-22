package repo_test

import (
	"quotes_api/internal/app/repositories"
	"testing"
)

func TestQuoteRepository_AddQuotes(t *testing.T) {
	repo := repositories.NewQuotesRepository()
	authorsData := []string{"A1", "A2", "A2", "A3"}
	quotesData := []string{"Q1", "Q2-1", "Q2-2", "Q3"}
	t.Run("Save", func(t *testing.T) {
		for i := range quotesData {
			if _, err := repo.Save(authorsData[i], quotesData[i]); err != nil {
				t.Error("repository cant save new data")
			}
		}
	})

	t.Run("GetAll", func(t *testing.T) {
		quotes, err := repo.GetAll()
		if err != nil {
			t.Errorf("Cant get all: %v", err)
		}
		if len(quotes) != 4 {
			t.Error("GetAll returned not all quotes")
		}
		if quotes[0].Author != "A1" || quotes[0].Content != "Q1" || quotes[0].Id != 1 ||
			quotes[1].Author != "A2" || quotes[1].Content != "Q2-1" || quotes[1].Id != 2 ||
			quotes[2].Author != "A2" || quotes[2].Content != "Q2-2" || quotes[2].Id != 3 ||
			quotes[3].Author != "A3" || quotes[3].Content != "Q3" || quotes[3].Id != 4 {
			t.Error("GetAll returned wrong data")
		}
	})

	t.Run("GetById", func(t *testing.T) {
		t.Run("Valid ID", func(t *testing.T) {
			quote, err := repo.GetById(1)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if quote == nil || quote.Author != "A1" || quote.Content != "Q1" {
				t.Error("Cant get correct author by ID")
			}
		})
		t.Run("Invalid ID", func(t *testing.T) {
			_, err := repo.GetById(0)
			if err == nil {
				t.Error("Expected error, but actually not happen because of invalid ID")
			}
		})
		t.Run("Value not exists", func(t *testing.T) {
			quote, err := repo.GetById(5)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if quote != nil {
				t.Errorf("Expected nil, actually got: %v", quote)
			}
		})
	})

	t.Run("GetByAuthor", func(t *testing.T) {
		t.Run("Existing author", func(t *testing.T) {
			quotes, err := repo.GetByAuthor("A2")
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if len(quotes) != 2 {
				t.Error("Failed to get all (2) authors by author name")
			}
			if quotes[0].Author != "A2" || quotes[0].Content != "Q2-1" || quotes[0].Id != 2 ||
				quotes[1].Author != "A2" || quotes[1].Content != "Q2-2" || quotes[1].Id != 3 {
				t.Error("Got wrong variables from author A2")
			}
		})
		t.Run("Non existing author", func(t *testing.T) {
			quotes, err := repo.GetByAuthor("A4")
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if len(quotes) != 0 {
				t.Errorf("Expected empty quotes cause of non-existing author name, actually : %v", quotes)
			}
		})
		t.Run("Empty author name", func(t *testing.T) {
			_, err := repo.GetByAuthor("")
			if err == nil {
				t.Error("Expected error cause empty author name, Actually nothing")
			}
		})
	})

	t.Run("Deletion", func(t *testing.T) {
		if err := repo.Delete(1); err != nil {
			t.Errorf("Unexpected deletion error: %v", err)
		}
		if err := repo.Delete(3); err != nil {
			t.Errorf("Unexpected deletion error: %v", err)
		}
		quotes, err := repo.GetAll()
		if err != nil {
			t.Errorf("Unexpected error after GetAll op inside Deletion test: %v", err)
		}
		if len(quotes) != 2 {
			t.Errorf("Expected len 2, got :%d", len(quotes))
		}
		if quotes[0].Id != 2 || quotes[1].Id != 4 ||
			quotes[0].Author != "A2" || quotes[1].Author != "A3" ||
			quotes[0].Content != "Q2-1" || quotes[1].Content != "Q3" {
			t.Errorf("Expected A2-Q2-1 and A3-Q3, actually got: %v", quotes)
		}
	})

}
