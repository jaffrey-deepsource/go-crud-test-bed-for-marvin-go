package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	httpapi "github.com/leeozebra/go-crud/internal/http"
	"github.com/leeozebra/go-crud/internal/repo"
	"github.com/leeozebra/go-crud/internal/service"
)

type healthResponse struct {
	Status string `json:"status"`
	Time   string `json:"time"`
}

func main() {

	memRepo := repo.NewBookRepoMem()
	bookSvc := service.NewBookService(memRepo)

	r := chi.NewRouter()
	r.Use(middleware.RequestID, middleware.RealIP, middleware.Logger, middleware.Recoverer)
	r.Use(jsonContentTypeMiddleware)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		resp := healthResponse{Status: "ok", Time: time.Now().UTC().Format(time.RFC3339Nano)}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(resp)
	})

	httpapi.MountBookRoutes(r, bookSvc)

	addr := ":8080"
	log.Printf("listening on http://localhost%v", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
