package helpers

import (
	"encoding/json"
	"net/http"
)

type baseResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type dataResponse[D any] struct {
	baseResponse
	Data D `json:"data"`
}

type errorResponse[E any] struct {
	baseResponse
	Error E `json:"error"`
}

func writeResponse(w http.ResponseWriter, payload any) {
	response, error := json.Marshal(payload)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(response)
}

func NewBaseResponse(
	w http.ResponseWriter,
	status int,
	message string,
) {
	w.Header().Add("Content-Type", "application/json")
	response := baseResponse{
		Status:  status,
		Message: message,
	}
	writeResponse(w, response)
}

func NewDataResponse(
	w http.ResponseWriter,
	status int,
	message string,
	data any,
) {
	w.Header().Add("Content-Type", "application/json")
	response := dataResponse[any]{
		baseResponse{
			Status:  status,
			Message: message,
		},
		data,
	}
	writeResponse(w, response)
}

func NewErrorResponse(
	w http.ResponseWriter,
	status int,
	message string,
	error any,
) {
	w.Header().Add("Content-Type", "application/json")
	response := errorResponse[any]{
		baseResponse{
			Status:  status,
			Message: message,
		},
		error,
	}
	writeResponse(w, response)
}
