package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rhiadc/betcontrol/config"
	"github.com/rhiadc/betcontrol/domain"
)

type Server struct {
	env     *config.Env
	Betfair domain.Betfair
}

func NewServer(Betfair domain.Betfair, env *config.Env) *Server {
	server := &Server{Betfair: Betfair, env: env}
	r := chi.NewRouter()
	server.router(r)
	s := &http.Server{Addr: env.APIAddr, Handler: r}
	log.Print(fmt.Sprintf("Server running on: %s", env.APIAddr))
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
	return server
}
