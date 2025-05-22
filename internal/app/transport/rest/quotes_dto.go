package rest

type CreateQuoteRequest struct {
	Author  string `json:"author"`
	Content string `json:"quote"`
}

type CreateQuoteResponse struct {
	Id      string `json:"id"`
	Author  string `json:"author"`
	Content string `json:"quote"`
}

type QuotesResponse struct {
	Quotes []SingleQuoteResponse `json:"quotes"`
}

type SingleQuoteResponse struct {
	Author  string `json:"author"`
	Content string `json:"quote"`
}
