package auto

import (
	"log"

	"tugasakhircoffe/TaCoffe/api/database"
	"tugasakhircoffe/TaCoffe/api/models"
	"tugasakhircoffe/TaCoffe/api/utils/console"
)

//Load to create database
func Load() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatal(err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatal(err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatal(err)
		}

		console.Pretty(users[i])
	}
}
