package domain

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	UserName string    `json:"user_name"`
	Password string    `json:"password"`
}

func (u *User) EncryptPassword() error {
	hasedPassword, err := hashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hasedPassword
	return nil
}

func (u *User) CheckPassword(password string) error {
	hasedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}
	if hasedPassword != u.Password {
		return ErrIncorrectPassword
	}
	return nil
}

func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}
