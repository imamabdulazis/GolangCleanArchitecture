package repository

import "tugasakhircoffe/TaCoffe/api/models"

type UserRepository interface {
	Save(models.User) (models.User, error)
	FindAll() ([]models.User, error)
	// 	FindByID(uint32, models.User) (int64, error)
	// 	Update(uint32, models.User) (int64, error)
	// 	Delete(uint32) int64
}
