package handler

import "github.com/gorilla/mux"

type handler struct {
	router *mux.Router
}

type HandlerOptions struct {
}

type Handler interface {
	GetRouter() *mux.Router
	RegisterRoutes()
}

func New(opts HandlerOptions) Handler {
	handler := &handler{
		router: mux.NewRouter(),
	}

	return handler
}

func (h *handler) GetRouter() *mux.Router {
	return h.router
}
