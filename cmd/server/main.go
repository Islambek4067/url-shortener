package main

import (
	"fmt"
	"log"
	"net/http"

	"shortener/internal/config"
	"shortener/internal/handler"
	"shortener/internal/storage"
)

func main() {
	store := storage.NewMemory()
	h := handler.NewHandler(store)
	cfg := config.Load()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "default")
	})
	mux.HandleFunc("POST /shorten", h.Shorten)
	mux.HandleFunc("GET /{code}", h.Redirect)
	mux.HandleFunc("GET /stats/{code}", h.Stats)

	log.Println("server on port :" + cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))

}
