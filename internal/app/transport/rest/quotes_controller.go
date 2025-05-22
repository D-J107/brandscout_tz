package rest

import (
	"encoding/json"
	"net/http"
	"quotes_api/internal/app/domain/models"
	"quotes_api/internal/app/services"
	"strconv"

	"github.com/gorilla/mux"
)

type QuoteController struct {
	qs services.QuotesService
}

func NewQuoteController() *QuoteController {
	return &QuoteController{qs: services.NewQuoteService()}
}

func (c *QuoteController) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateQuoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "invalid request format body")
		return
	}

	created, err := c.qs.Save(req.Author, req.Content)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	resp := CreateQuoteResponse{
		Id:      strconv.Itoa(created.Id),
		Author:  created.Author,
		Content: created.Content,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (c *QuoteController) GetByFilter(w http.ResponseWriter, r *http.Request) {
	var quotes []*models.Quote
	author := r.URL.Query().Get("author")

	if author == "" { // получение всех цитат без фильтрации по автору
		var err error
		quotes, err = c.qs.GetAll()
		if err != nil {
			WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		var err error
		quotes, err = c.qs.GetByAuthor(author)
		if err != nil {
			WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	resp := QuotesResponse{
		Quotes: make([]SingleQuoteResponse, 0),
	}
	for _, quote := range quotes {
		resp.Quotes = append(resp.Quotes, SingleQuoteResponse{
			Author:  quote.Author,
			Content: quote.Content,
		})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (c *QuoteController) GetRandom(w http.ResponseWriter, r *http.Request) {
	quote, err := c.qs.GetRandom()
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := SingleQuoteResponse{
		Author:  quote.Author,
		Content: quote.Content,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (c *QuoteController) DeleteById(w http.ResponseWriter, r *http.Request) {
	quoteId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "invalid id integer format")
		return
	}

	if err := c.qs.DeleteById(quoteId); err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
