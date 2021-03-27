package controller

import (
	"net/http"
	"tugasakhircoffe/TaCoffe/api/database"
	"tugasakhircoffe/TaCoffe/api/repository"
	"tugasakhircoffe/TaCoffe/api/repository/crud"
	"tugasakhircoffe/TaCoffe/api/responses"
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
		responses.JSON(w, http.StatusOK, "ok", users)
	}(repo)
}
