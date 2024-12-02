package service

import (
	"encoding/json"
	"gateway/risk/model"
	"net/http"
	"time"
)

// Exported function to handle requests
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	time.Sleep(50 * time.Millisecond)
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w)
	} else {
		resp := model.JSONResponse{
			Result:  "OK",
			Service: "Risk",
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	err := FindCookie(w, r)
	if !err {
		errResp := model.JSONResponse{
			Result:  "Not Authorized",
			Service: "Please use '/' first to get cookies",
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errResp)
		return
	}

	response := map[string]string{
		"Username": "arga",
		"Name":     "Arga",
		"Address":  "Jalan Sudirman",
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
