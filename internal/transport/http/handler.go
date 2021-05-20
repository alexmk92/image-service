package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/alexmk92/image-service/internal/image"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Stores the pointer to the stock service
type Handler struct{
    Router *mux.Router
    Service *image.Service
}

// Response - an object to store responses from our API
type Response struct {
    Message string
    Error string
}

// Returns a pointer to a Handler
func NewHandler(service *image.Service) *Handler {
    return &Handler{
        Service: service,
    }
}

// BasicAuth
func BasicAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Info("Trying to authenticate user")
        user, pass, ok := r.BasicAuth()
        if user == "admin" && pass == "password" && ok {
            original(w, r)
        } else {
            w.Header().Set("Content-Type", "application/json; charset=utf-8")
            sendErrorResponse(w, "not authorized", errors.New("not authorized"))
        }
    }
}

func validateToken(accessToken string) bool {
    var signingKey = []byte("s3542edgwg4#@24'-3434sgsd2")
    token, err := jwt.Parse(accessToken, func (token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("error signing token")
        }

        return signingKey, nil
    })

    if err != nil {
        return false
    }

    return token.Valid
}

// JWTAuth - a decorator for JWT validation on endpoints
func JWTAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Info("Trying to authenticate JWT")
        authHeader := r.Header["Authorization"]
        if authHeader == nil {
            w.Header().Set("Content-Type", "application/json; charset=utf-8")
            sendErrorResponse(w, "not authorized", errors.New("not authorized, not authorization header set"))
        }

        // Bearer token
        authHeaderParts := strings.Split(authHeader[0], " ")
        if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
            w.Header().Set("Content-Type", "application/json; charset=utf-8")
            sendErrorResponse(w, "not authorized", errors.New("not authorized, expected Bearer TOKEN"))
        }

        if validateToken(authHeaderParts[1]) {
            original(w, r)
        } else {
            w.Header().Set("Content-Type", "application/json; charset=utf-8")
            sendErrorResponse(w, "not authorized", errors.New("not authorized"))
        }
    }
}

// LoggingMiddleware - adds middleware around endpoints.
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.WithFields(
            log.Fields{
                "Method": r.Method,
                "Path": r.URL.Path,
            },
        ).Info("handled request")
        next.ServeHTTP(w, r)
    })
}

// SetupRoutes - sets up all routes for the application
func (h *Handler) SetupRoutes() {
    log.Info("Setting up routes")
    h.Router = mux.NewRouter()
    h.Router.Use(LoggingMiddleware)
    h.Router.HandleFunc("/api/health", func (w http.ResponseWriter, _ *http.Request){
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode(Response{Message: "I am alive!"}); err != nil {
            panic(err)
        }
    })
    h.Router.HandleFunc("/api/image/{id}", h.GetImage).Methods("GET")
    h.Router.HandleFunc("/api/image/{id}", BasicAuth(h.UpdateImage)).Methods("PUT")
    h.Router.HandleFunc("/api/image", BasicAuth(h.CreateImage)).Methods("POST")
    h.Router.HandleFunc("/api/image/{id}", JWTAuth(h.DeleteImage)).Methods("DELETE")
    h.Router.HandleFunc("/api/image", JWTAuth(h.GetAllImages)).Methods("GET")
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
    w.WriteHeader(http.StatusInternalServerError)
    if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
        panic(err)
    }
}
