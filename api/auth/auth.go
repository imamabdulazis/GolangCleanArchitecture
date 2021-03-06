package auth

import (
	"tugasakhircoffe/TaCoffe/api/database"
	"tugasakhircoffe/TaCoffe/api/models"
	"tugasakhircoffe/TaCoffe/api/security"
	"tugasakhircoffe/TaCoffe/api/utils/channels"

	"github.com/jinzhu/gorm"
)

// Login method
func Login(username, password string) (string, error) {
	user := models.User{}
	var err error
	var db *gorm.DB
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		db, err = database.ConnectDB()
		if err != nil {
			ch <- false
			return
		}
		defer db.Close()

		err = db.Debug().Model(models.User{}).Where("username = ?", username).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}

		err = security.VerifyPassword(user.Password, password)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		user.Password = ""
		return GenerateJWT(user)
	}

	return "", err
}
