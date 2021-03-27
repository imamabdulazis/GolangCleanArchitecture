package models

import (
	"errors"
	"html"
	"strings"
	"tugasakhircoffe/TaCoffe/api/security"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
)

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

// BeforeSave hash the user password
func (u *User) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Prepare cleans the inputs
func (u *User) Prepare() {
	u.ID = 0
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
}

// Validate validates the inputs
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "login":
		if u.Username == "" {
			return errors.New("Username is required")
		}
		if u.Password == "" {
			return errors.New("Password is required")
		}
	default:
		if u.Username == "" {
			return errors.New("Username is required")
		}

		if u.Password == "" {
			return errors.New("Password is required")
		}

		if u.Email == "" {
			return errors.New("Email is required")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email address")
		}
	}
	return nil
}
