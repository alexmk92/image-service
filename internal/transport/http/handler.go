package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Stores the pointer to the stock service
type Handler struct{
    Router *mux.Router
}

// Returns a pointer to a Handler
func NewHandler() *Handler {
    return &Handler{}
}

// SetupRoutes - sets up all routes for the application
func (h *Handler) SetupRoutes() {
    fmt.Println("Setting up routes")
    h.Router = mux.NewRouter()
    h.Router.HandleFunc("/api/health", func (w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Alive!")
    })
}
