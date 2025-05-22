package app

import (
	"fmt"
	"net/http"
	"quotes_api/internal/app/config"
)

func RunServer() {

	cfg := config.MustLoad()

	router := setupRoutes()
	go func() {
		fmt.Println("REST API running on " + cfg.RestPort)
		if err := http.ListenAndServe(cfg.RestPort, router); err != nil {
			fmt.Println("REST API quotes server error: ", err.Error())
		}
	}()

	// бесконечное ожидание чтобы приложение не завершилось и в горутинах обратабывались события
	select {}
}
