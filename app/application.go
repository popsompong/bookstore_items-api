package app

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapUrls()

	srv := &http.Server{
		Addr:              "127.0.0.1:8080",
		WriteTimeout:      500 * time.Millisecond,
		ReadHeaderTimeout: 2 * time.Second,
		IdleTimeout:       60 * time.Second,
		Handler:           router,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

}
