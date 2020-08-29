package app

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/shakilbd009/go-items-api/src/clients/elasticsearch"
)

const (
	location = "127.0.0.1:8083"
)

var (
	router = mux.NewRouter()
)

func StartApp() {
	elasticsearch.Init()
	urlMapping()

	srv := &http.Server{
		Addr:         location,
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
