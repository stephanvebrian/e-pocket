package handler

import "net/http"

func (h *handler) Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}
