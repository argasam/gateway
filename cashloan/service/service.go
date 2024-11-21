package service

import (
	"encoding/json"
	"gateway/cashloan/model"
	"net/http"
)

// Exported function to handle requests
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w)
	} else {
		resp := model.JSONResponse{
			Result: "OK",
			Reason: "Cashloan Service",
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
