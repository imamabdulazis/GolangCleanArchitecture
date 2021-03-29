package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"tugasakhircoffe/TaCoffe/api/security"

	"github.com/badoux/checkmail"

	_ "github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//User models
type User struct {
	ID         uuid.UUID  `gorm:"primary_key; unique; type:uuid;column:id;default:uuid_generate_v4()" json:"id"`
	Username   string     `json:"username"`
	Email      string     `json:"email"`
	Name       string     `json:"name"`
	Password   string     `json:"password,omitempty"`
	ImageUrl   string     `json:"image_url"`
	TelpNumber string     `json:"telp_number"`
	Role       string     `json:"role"`
	Address    string     `json:"address"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `sql:"index" json:"-"`
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

		if u.Name == "" {
			return errors.New("Name is required")
		}

		if u.Role == "" {
			return errors.New("Role is requred")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email address")
		}
	}
	return nil
}
