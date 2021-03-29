package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"tugasakhircoffe/TaCoffe/api/database"
	"tugasakhircoffe/TaCoffe/api/models"
	"tugasakhircoffe/TaCoffe/api/repository"
	"tugasakhircoffe/TaCoffe/api/repository/crud"
	"tugasakhircoffe/TaCoffe/api/responses"
	"tugasakhircoffe/TaCoffe/helper"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

// GET All users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectDB()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryUsersCRUD(db)

	func(userRepository repository.UserRepository) {
		users, err := userRepository.FindAll()
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusOK, "Ok", users)
	}(repo)
}

//Create user
func CreateUsers(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		responses.ERROR(w, http.StatusConflict, err)
		return
	}

	//validate user value
	user.Prepare()
	err = user.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//connecting to db
	db, err := database.ConnectDB()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryUsersCRUD(db)

	func(userRepository repository.UserRepository) {
		user, err := userRepository.Save(user)
		if err != nil {
			responses.ERROR(w, http.StatusConflict, err)
			return
		}

		w.Header().Set("Location,", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user.ID))
		responses.JSON(w, http.StatusCreated, "Ok", user)
	}(repo)

}

//Get details user
func GetUserDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := uuid.FromString(vars["id"])

	if helper.IsValidUUID(uid.String()) == false {
		responses.JSON_ERROR(w, http.StatusBadRequest, "Id yang anda masukan tidak valid")
		return
	}
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryUsersCRUD(db)

	func(userRepository repository.UserRepository) {
		user, err := userRepository.FindByID(uid)
		if err != nil {
			if err.Error() == "record not found" {
				responses.JSON_ERROR(w, http.StatusBadRequest, "User tidak ditemukan")
				return
			}
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		responses.JSON(w, http.StatusOK, "Ok", user)
	}(repo)
}

//Update user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := uuid.FromString(vars["id"])

	if helper.IsValidUUID(uid.String()) == false {
		responses.JSON_ERROR(w, http.StatusBadRequest, "Id yang anda masukan tidak valid")
		return
	}
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user := models.User{}
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryUsersCRUD(db)

	func(userRepository repository.UserRepository) {
		_, err := userRepository.Update(uid, user)
		if err != nil {
			if err.Error() == "record not found" {
				responses.JSON_ERROR(w, http.StatusBadRequest, "User tidak ditemukan")
				return
			}
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		responses.JSON_SUCCESS(w, http.StatusOK, "Berhasil update user")
	}(repo)
}

//Delete user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := uuid.FromString(vars["id"])

	if helper.IsValidUUID(uid.String()) == false {
		responses.JSON_ERROR(w, http.StatusBadRequest, "Id yang anda masukan tidak valid")
		return
	}
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	if helper.IsValidUUID(uid.String()) == false {
		responses.JSON_ERROR(w, http.StatusUnprocessableEntity, "Id yang anda masukan tidak valid")
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := crud.NewRepositoryUsersCRUD(db)

	func(userRepository repository.UserRepository) {
		_, err = userRepository.Delete(uid)
		if err != nil {
			if err.Error() == "record not found" {
				responses.JSON_ERROR(w, http.StatusBadRequest, "User tidak ditemukan")
				return
			}
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		w.Header().Set("Entity", fmt.Sprintf("%d", uid))
		responses.JSON_SUCCESS(w, http.StatusOK, "Hapus berhasil")
	}(repo)

}
