package controllers

import (
	"encoding/json"
	"net/http"

	"tugasakhircoffe/TaCoffe/api/auth"
	"tugasakhircoffe/TaCoffe/api/models"
	"tugasakhircoffe/TaCoffe/api/responses"
)

func Login(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.Login(user.Username, user.Password)
	if err != nil {
		if err.Error() == "record not found" {
			responses.JSON_ERROR(w, http.StatusUnauthorized, "Username belum terdaftar")
			return
		}
		responses.JSON_ERROR(w, http.StatusUnauthorized, "Password anda tidak valid")
		return
	}
	responses.AUTH_JSON(w, http.StatusOK, "Ok", token)

}
