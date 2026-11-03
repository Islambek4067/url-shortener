package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"shortener/internal/storage"
)

type Handler struct {
	store storage.Storage
}

func NewHandler(store storage.Storage) *Handler {
	return &Handler{
		store: store,
	}
}

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	ShortURL string `json:"short_url"`
}

func (h *Handler) Shorten(w http.ResponseWriter, r *http.Request) {
	var req shortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "not correct JSON", http.StatusBadRequest)
		return
	}

	u, err := url.ParseRequestURI(req.URL)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") {
		http.Error(w, "not correct URL", http.StatusBadRequest)
		return
	}

	code := storage.GenerateCode(6)
	if err := h.store.Save(code, req.URL); err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Contenet-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(shortenResponse{
		ShortURL: "http://" + r.Host + "/" + code,
	})
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	link, ok, err := h.store.Get(code)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	if !ok {

		http.Error(w, "This link doesn't exist", http.StatusNotFound)
		return
	}

	if err := h.store.AddClick(code); err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, link.Original, http.StatusFound)
}

type statsResponse struct {
	Original  string    `json:"original"`
	Clicks    int       `json:"clicks"`
	CreatedAt time.Time `json:"created_at"`
}

func (h *Handler) Stats(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	link, ok, err := h.store.Get(code)

	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(statsResponse{
		Original:  link.Original,
		Clicks:    link.Clicks,
		CreatedAt: link.CreatedAt,
	})
}
