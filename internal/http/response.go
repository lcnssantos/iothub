package http

import (
	"encoding/json"
	"net/http"
)

func SetResponse(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	if data != nil {
		answer, _ := json.Marshal(data)
		w.Write(answer)
	}
}
