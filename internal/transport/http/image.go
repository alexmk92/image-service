package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/alexmk92/image-service/internal/image"
	"github.com/gorilla/mux"
)

// GetImage - retrieve an image by its ID
func (h *Handler) GetImage(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    vars := mux.Vars(r)
    id := vars["id"]

    i, err := strconv.ParseUint(id, 10, 64)
    if err != nil {
        sendErrorResponse(w, "Unable to parse UINT from id", err)
    }

    image, err := h.Service.GetImage(uint(i))
    if json.NewEncoder(w).Encode(image); err != nil {
        panic(err)
    }
}

// UpdateImage - updates an image
func (h *Handler) UpdateImage(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusOK)

    var image image.Image
    if err := json.NewDecoder(r.Body).Decode(&image); err != nil {
        sendErrorResponse(w, "Failed to decode JSON body", err)
    }

    vars := mux.Vars(r)
    id := vars["id"]
    imageId, err := strconv.ParseUint(id, 10, 64)
    if err != nil {
        sendErrorResponse(w, "Unable to parse UINT from id", err)
    }

    image, err = h.Service.UpdateImage(uint(imageId), image)
    if json.NewEncoder(w).Encode(image); err != nil {
        panic(err)
    }
}

// UpdateImage - updates an image
func (h *Handler) DeleteImage(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusOK)

    vars := mux.Vars(r)
    id := vars["id"]
    imageID, err := strconv.ParseUint(id, 10, 64)

    if err != nil {
        sendErrorResponse(w, "Failed to parse uint from ID", err)
    }

    err = h.Service.DeleteImage(uint(imageID))
    if err != nil {
        json.NewEncoder(w).Encode(Response{Message:"Failed to delete image by ID"});
        panic(err)
    }
}

// PostImage - updates an image
func (h *Handler) CreateImage(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusOK)

    var image image.Image
    if err := json.NewDecoder(r.Body).Decode(&image); err != nil {
        sendErrorResponse(w, "Failed to decode JSON body", err)
    }

    image, err := h.Service.PostImage(image)

    if json.NewEncoder(w).Encode(image); err != nil {
        panic(err)
    }
}

// GetAllImages - retrieve images from the service
func (h *Handler) GetAllImages(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    _, err := h.Service.GetAllImages()
    if err != nil {
        sendErrorResponse(w, "Failed to retrieve all images", err)
    }
}
