package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Handler stores pointer to comments service
type Handler struct {
	Router *mux.Router
}

// NewHandler returns a pointer to Handler
func NewHandler() *Handler {
	return &Handler{}
}

// SetupRoutes sets up routing for the application
func (h *Handler) SetupRoutes() {
	fmt.Println("Initializing routes...")

	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})
}
