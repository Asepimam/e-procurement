package response

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // bisa nil untuk error
	Token  string      	`json:"token,omitempty"` // optional token field for authentication responses
	Meta 	*Meta		`json:"meta,omitempty"`
}

// struct for pagination metadata
type Meta struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	HasMore    bool `json:"has_more"` // indicates if there are more pages
}

// function for sending JSON responses mapping to the ApiResponse struct
func Success(w http.ResponseWriter, message string, data interface{}, meta *Meta, token ...string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    var tokenStr string
    if len(token) > 0 {
        tokenStr = token[0]
    }

    json.NewEncoder(w).Encode(ApiResponse{
        Status:  "success",
        Message: message,
        Data:    data,
        Token:   tokenStr,
        Meta:    meta,
    })
}

func Error(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ApiResponse{
		Status:  "error",
		Message: message,
		Data:    nil, // null for error responses
		Meta:    nil, // 
	})
}
