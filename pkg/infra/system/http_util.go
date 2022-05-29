package system

import (
	json "encoding/json"
	http "net/http"
)

// HTTPResponseWithCode - Set code on http response
func HTTPResponseWithCode(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}

// HTTPResponseWithJSON - Set code and json on http response
func HTTPResponseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	w.Write(response)
}

// HTTPResponseWithError - Set code and message error on http response
func HTTPResponseWithError(w http.ResponseWriter, code int, message string) {
	HTTPResponseWithJSON(w, code, map[string]string{"error": message})
}
