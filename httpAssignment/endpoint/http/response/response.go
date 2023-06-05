package response

import "net/http"

func SendOK(w http.ResponseWriter, data any) {
	sendJson(w, http.StatusOK, data)
}
