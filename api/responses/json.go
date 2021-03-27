package responses

import (
	"encoding/json"
	"net/http"
)

// auth struct
type AuthResponseHandler struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

//response struct
type ResponseHandler struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//response struct
type ResponseHandlerError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

//JSON AUTH displays 1 a json response message with data
func AUTH_JSON(w http.ResponseWriter, statusCode int, message string, token string) {
	w.WriteHeader(statusCode)
	var value = AuthResponseHandler{
		Status:  statusCode,
		Message: message,
		Token:   token,
	}
	err := json.NewEncoder(w).Encode(value)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

// JSON displays a json response message with data
func JSON(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.WriteHeader(statusCode)
	var value = ResponseHandler{
		Status:  statusCode,
		Message: message,
		Data:    data,
	}
	err := json.NewEncoder(w).Encode(value)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

// JSON displays a json response message with data
func JSON_ERROR(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	var value = ResponseHandlerError{
		Status:  statusCode,
		Message: message,
	}
	err := json.NewEncoder(w).Encode(value)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		if statusCode == 400 {
			JSON_ERROR(w, http.StatusBadRequest, err.Error())
		}
		JSON_ERROR(w, http.StatusUnauthorized, err.Error())
		return
	}
	JSON_ERROR(w, http.StatusBadRequest, "Terjadi kesalahan, mohon coba kembali")
}
