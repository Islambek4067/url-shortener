package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"shortener/internal/config"
	"shortener/internal/handler"
	"shortener/internal/storage"
)

func main() {

	cfg := config.Load()
	var store storage.Storage

	if cfg.DatabaseURL != "" {
		pg, err := storage.NewPostgres(context.Background(), cfg.DatabaseURL)
		if err != nil {
			log.Fatal(err)
		}
		store = pg
		log.Println("storage: PostgreSQL")
	} else {
		store = storage.NewMemory()
		log.Println("storage: memory")
	}

	h := handler.NewHandler(store)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /default", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "default")
	})
	mux.HandleFunc("POST /shorten", h.Shorten)
	mux.HandleFunc("GET /{code}", h.Redirect)
	mux.HandleFunc("GET /stats/{code}", h.Stats)

	log.Println("server on port :" + cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))

}
