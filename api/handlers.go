package api

import (
	"fmt"
	"net/http"

	"github.com/rhiadc/betcontrol/domain"
)

type Handler struct {
	service domain.Betfair
}

func NewHandler(service domain.Betfair) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CallExternalAPI(w http.ResponseWriter, r *http.Request) {
	value := h.service.CallBetfairAPI()
	w.Write([]byte(fmt.Sprintf(value)))
}
