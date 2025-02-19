package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/mux"
	"github.com/stephanvebrian/e-pocket/pocket-engine/logic/transfer"
	"github.com/stephanvebrian/e-pocket/pocket-engine/middleware"
	"github.com/stephanvebrian/e-pocket/pocket-engine/model"
)

type handler struct {
	router        *mux.Router
	transferLogic transfer.TransferLogic
}

type HandlerOptions struct {
	TransferLogic transfer.TransferLogic
}

type Handler interface {
	GetRouter() *mux.Router
	RegisterRoutes()
}

func New(opts HandlerOptions) Handler {
	handler := &handler{
		transferLogic: opts.TransferLogic,
	}

	return handler
}

func (h *handler) GetRouter() *mux.Router {
	return h.router
}

func (h *handler) RegisterRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/", h.Hello)

	router.Use(middleware.StartLoggingMiddleware)

	// transfers
	router.HandleFunc("/v1/transfer", h.CreateTransfer).Methods("POST")

	router.Use(middleware.EndLoggingMiddleware)

	fmt.Println("Routes registered")

	h.router = router
}

func (h *handler) writeError(ctx context.Context, opt model.ErrorResponseOption) {
	opt.Response.Timestamp = time.Now().UTC().Format(time.RFC3339) // Add timestamp
	opt.Writer.Header().Set("Content-Type", "application/json")
	opt.Writer.WriteHeader(opt.StatusCode)
	json.NewEncoder(opt.Writer).Encode(opt.Response)
}
