package handler

import "fmt"

func (h *handler) RegisterRoutes() {
	h.router.HandleFunc("/", h.Hello)

	fmt.Println("Routes registered")
}
