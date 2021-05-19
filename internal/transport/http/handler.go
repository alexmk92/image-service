package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alexmk92/image-service/internal/image"
	"github.com/gorilla/mux"
)

// Stores the pointer to the stock service
type Handler struct{
    Router *mux.Router
    Service *image.Service
}

// Returns a pointer to a Handler
func NewHandler(service *image.Service) *Handler {
    return &Handler{
        Service: service,
    }
}

// SetupRoutes - sets up all routes for the application
func (h *Handler) SetupRoutes() {
    fmt.Println("Setting up routes")
    h.Router = mux.NewRouter()
    h.Router.HandleFunc("/api/health", func (w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Alive!")
    })
    h.Router.HandleFunc("/api/image/{id}", h.GetImage).Methods("GET")
    h.Router.HandleFunc("/api/image/{id}", h.UpdateImage).Methods("PUT")
    h.Router.HandleFunc("/api/image", h.CreateImage).Methods("POST")
    h.Router.HandleFunc("/api/image/{id}", h.DeleteImage).Methods("DELETE")
    h.Router.HandleFunc("/api/image", h.GetAllImages).Methods("GET")
}

// GetImage - retrieve an image by its ID
func (h *Handler) GetImage(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    i, err := strconv.ParseUint(id, 10, 64)
    if err != nil {
        fmt.Fprintf(w, "Unable to parse UINT from id")
    }

    image, err := h.Service.GetImage(uint(i))
    if err != nil {
        fmt.Fprintf(w, "Error retrieving image by ID")
    }

    fmt.Fprintf(w, "%+v", image)
}

// UpdateImage - updates an image
func (h *Handler) UpdateImage(w http.ResponseWriter, r *http.Request) {
    image, err := h.Service.UpdateImage(1, image.Image{
        StorageUrl: "/new",
    })

    if err != nil {
        fmt.Fprintf(w, "Failed to update image")
    }

    fmt.Fprintf(w, "%+v", image)
}

// UpdateImage - updates an image
func (h *Handler) DeleteImage(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    imageID, err := strconv.ParseUint(id, 10, 64)
    if err != nil {
        fmt.Fprintf(w, "Failed to parse uint from ID")
    }

    err = h.Service.DeleteImage(uint(imageID))
    if err != nil {
        fmt.Fprintf(w, "Failed to delete image by image ID")
    }

    fmt.Fprintf(w, "Successfully deleted image")
}

// PostImage - updates an image
func (h *Handler) CreateImage(w http.ResponseWriter, r *http.Request) {
    image, err := h.Service.PostImage(image.Image{
        StorageUrl: "/",
        AspectRatio: "1:1",
    })

    if err != nil {
        fmt.Fprintf(w, "Failed to post image")
    }

    fmt.Fprintf(w, "%+v", image)
}

// GetAllImages - retrieve images from the service
func (h *Handler) GetAllImages(w http.ResponseWriter, r *http.Request) {
    images, err := h.Service.GetAllImages()
    if err != nil {
        fmt.Fprintf(w, "Failed to retrieve all images")
    }

    fmt.Fprintf(w, "%+v", images)
}
