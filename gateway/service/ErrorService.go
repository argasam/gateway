package service

import (
	"encoding/json"
	"gateway/model"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, StatusCode int, message string) {
	errResp := model.JSONResponse{Code: StatusCode, Result: message}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(StatusCode)
	json.NewEncoder(w).Encode(errResp)
}
