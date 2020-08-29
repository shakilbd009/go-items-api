package app

import (
	"net/http"

	"github.com/shakilbd009/go-items-api/src/controllers"
)

func urlMapping() {

	router.HandleFunc("/ping", controllers.PingController.Get).Methods(http.MethodGet)
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", controllers.ItemsController.Get).Methods(http.MethodGet)
	router.HandleFunc("/items/{id}", controllers.ItemsController.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/items/search", controllers.ItemsController.Search).Methods(http.MethodPost)
}
