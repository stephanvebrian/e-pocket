package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/stephanvebrian/e-pocket/pocket-engine/handler"
)

func main() {
	handler := handler.New(handler.HandlerOptions{})
	handler.RegisterRoutes()

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":3010", handler.GetRouter()))

	fmt.Println("Server started on port 3010")
}
