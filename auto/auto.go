package auto

import (
	"tugasakhircoffe/TaCoffe/api/models"

	"github.com/google/uuid"
)

var users = []models.User{
	models.User{
		ID:         uuid.UUID{},
		Username:   "sample",
		Email:      "sample@gmail.com",
		Name:       "sample",
		Password:   "12345",
		ImageUrl:   "url",
		TelpNumber: "081",
		Role:       1,
		Address:    "sample_address",
	},
}
