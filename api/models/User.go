package models

import "github.com/jinzhu/gorm"

//User models
type User struct {
	gorm.Model
	Username   string
	Email      string
	Name       string
	Password   string
	ImageUrl   string
	TelpNumber string
	Role       int
	Address    string
}
