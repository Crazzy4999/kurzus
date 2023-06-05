package response

import (
	"encoding/json"
	"net/http"
)

func sendJson(w http.ResponseWriter, code int, data any) {
	jsonData, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonData)
}
