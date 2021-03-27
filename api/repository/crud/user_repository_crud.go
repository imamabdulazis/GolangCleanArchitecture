package crud

import (
	"tugasakhircoffe/TaCoffe/api/models"
	"tugasakhircoffe/TaCoffe/api/utils/channels"

	"github.com/jinzhu/gorm"
)

type RepositoryUsersCRUD struct {
	db *gorm.DB
}

// NewRepositoryUsersCRUD returns a new repository with DB connection
func NewRepositoryUsersCRUD(db *gorm.DB) *RepositoryUsersCRUD {
	return &RepositoryUsersCRUD{db}
}

// Save returns a new user created or an error
func (r *RepositoryUsersCRUD) Save(user models.User) (models.User, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return user, nil
	}
	return models.User{}, err
}

//FindAll
func (r *RepositoryUsersCRUD) FindAll() ([]models.User, error) {
	var err error
	users := []models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Limit(100).Find(&users).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return users, nil
	}
	return nil, err
}
