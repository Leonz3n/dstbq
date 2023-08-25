package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HttpServer struct {
	router *httprouter.Router
}

func (h *HttpServer) Start() error {
	h.router = httprouter.New()

	http.ListenAndServe(":8080", h.router)

	return nil
}

func (h *HttpServer) Shutdown() {

}
