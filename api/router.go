package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func (s *Server) router(r *chi.Mux) {
	handler := NewHandler(s.Betfair)
	r.Get("/", handler.CallExternalAPI)
	r.Get("/slow", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		w.Write([]byte(fmt.Sprintf("All done")))
	})
}
