package repository

import (
	"tugasakhircoffe/TaCoffe/api/models"

	uuid "github.com/satori/go.uuid"
)

type UserRepository interface {
	Save(models.User) (models.User, error)
	FindAll() ([]models.User, error)
	FindByID(uuid.UUID) (models.User, error)
	Update(uuid.UUID, models.User) (int64, error)
	Delete(uuid.UUID) (int64, error)
}
