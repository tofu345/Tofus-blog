package routes

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	ResponseCode int    `json:"responseCode"`
	Message      string `json:"message"`
	Data         any    `json:"data"`
}

func JSONResponse(w http.ResponseWriter, responseCode int, data any, message string) {
	w.Header().Set("Content-type", "application/json")

	if responseCode == 103 {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(Response{
		ResponseCode: responseCode,
		Data:         data,
		Message:      message,
	})
}
