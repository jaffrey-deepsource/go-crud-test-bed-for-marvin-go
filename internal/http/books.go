package httpapi

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/leeozebra/go-crud/internal/domain"
	"github.com/leeozebra/go-crud/internal/repo"
	"github.com/leeozebra/go-crud/internal/service"
)

type BookHandlers struct {
	svc *service.BookService
}

func MountBookRoutes(r *chi.Mux, svc *service.BookService) {
	h := &BookHandlers{svc: svc}
	r.Route("/v1/books", func(r chi.Router) {
		r.Post("/", h.create)
		r.Get("/", h.list)
		r.Get("/{id}", h.getByID)
		r.Patch("/{id}", h.updatePartial)
		r.Delete("/{id}", h.delete)
	})
}

func (h *BookHandlers) create(w http.ResponseWriter, r *http.Request) {
	var in domain.CreateBookInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, `{"error":"invalid body"}`, http.StatusBadRequest)
		return
	}
	b, err := h.svc.Create(r.Context(), in)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(b)
}

func (h *BookHandlers) getByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	b, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		status := http.StatusInternalServerError
		if err == repo.ErrNotFound {
			status = http.StatusNotFound
		}
		http.Error(w, `{"error":"not found"}`, status)
		return
	}
	_ = json.NewEncoder(w).Encode(b)
}

func (h *BookHandlers) list(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	items, err := h.svc.List(r.Context(), limit, offset)
	if err != nil {
		http.Error(w, `{"error":"bad request"}`, http.StatusBadRequest)
		return
	}
	_ = json.NewEncoder(w).Encode(struct {
		Items []domain.Book `json:"items"`
	}{Items: items})
}

func (h *BookHandlers) updatePartial(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var in domain.UpdateBookInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, `{"error":"invalid body"}`, http.StatusBadRequest)
		return
	}
	b, err := h.svc.UpdatePartial(r.Context(), id, in)
	if err != nil {
		status := http.StatusInternalServerError
		if err == repo.ErrNotFound {
			status = http.StatusNotFound
		}
		http.Error(w, `{"error":"not found"}`, status)
		return
	}
	_ = json.NewEncoder(w).Encode(b)
}

func (h *BookHandlers) delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.svc.Delete(r.Context(), id); err != nil {
		status := http.StatusInternalServerError
		if err == repo.ErrNotFound {
			status = http.StatusNotFound
		}
		http.Error(w, `{"error":"not found"}`, status)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
