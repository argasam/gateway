package service

import (
	"encoding/json"
	"gateway/model"
	"io"
	"log"
	"net/http"
)

func ServiceRequest(w http.ResponseWriter, r *http.Request, serviceURL string) {
	req, err := http.NewRequest(r.Method, serviceURL, r.Body)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Failed to create request")
		return
	}

	// Copy headers from the original request
	req.Header = r.Header

	// Send the request to the service
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ErrorResponse(w, http.StatusNotImplemented, "Gateway is available, but downstream services are unavailable")
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Gateway internal exception")
		return
	}

	// Parse the service response
	var serviceResponse map[string]interface{}
	if err := json.Unmarshal(body, &serviceResponse); err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Gateway internal exception")
		return
	}
	log.Println(serviceResponse)

	// Create the encapsulated response
	response := model.JSONResponse{
		Code:   resp.StatusCode,
		Result: getResponse(serviceResponse),
	}

	cookies := resp.Cookies()
	log.Println(cookies)
	log.Println(len(cookies))
	if len(cookies) > 0 {
		for _, cookie := range cookies {
			http.SetCookie(w, cookie)
			log.Printf("%s %s %s %v", cookie.Name, cookie.Value, cookie.Path, cookie.HttpOnly)
		}
	}

	// Set the header and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(response)
}

func getResponse(rawResponse map[string]interface{}) interface{} {
	if result, isThere := rawResponse["Result"]; result != "" && isThere {
		return result
	}
	return rawResponse
}
