package chi_middlewar

import (
	"net/http"
	"strings"
)

func JSONContentTypeMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Skip chanking for GET and DELETE requests
        if r.Method == http.MethodGet || r.Method == http.MethodDelete {
            next.ServeHTTP(w, r)
            return
        }
        
        // Check Content-Type header
        contentType := r.Header.Get("Content-Type")
        if contentType == "" {
            http.Error(w, "Content-Type header is required", http.StatusBadRequest)
            return
        }
        
        // Check if Content-Type is application/json
        if !strings.HasPrefix(contentType, "application/json") {
            http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
            return
        }
        
        next.ServeHTTP(w, r)
    })
}