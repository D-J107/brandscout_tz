package app

import (
	"quotes_api/internal/app/transport/rest"

	"github.com/gorilla/mux"
)

func setupRoutes() *mux.Router {
	quotesController := rest.NewQuoteController()

	router := mux.NewRouter()

	router.HandleFunc("/quotes", quotesController.Create).Methods("POST")
	router.HandleFunc("/quotes", quotesController.GetByFilter).Methods("GET")
	router.HandleFunc("/quotes/random", quotesController.GetRandom).Methods("GET")
	router.HandleFunc("/quotes/{id}", quotesController.DeleteById).Methods("DELETE")

	return router
}
