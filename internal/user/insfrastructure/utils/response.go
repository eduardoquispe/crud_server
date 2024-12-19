package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func JSON(w http.ResponseWriter, statusCode int, success bool, message string, data interface{}) { 
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	res := Response{
		Success: success,
		Message: message,
		Data:    data,
	}
	
	json.NewEncoder(w).Encode(res)
}