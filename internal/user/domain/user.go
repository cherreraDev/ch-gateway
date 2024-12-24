package domain

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id       uuid.UUID
	userName string
	password string
}

func (u *User) EncryptPassword() error {
	hasedPassword, err := hashPassword(u.password)
	if err != nil {
		return err
	}
	u.password = hasedPassword
	return nil
}

func (u *User) CheckPassword(password string) error {
	hasedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}
	if hasedPassword != u.password {
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

func (u *User) Id() uuid.UUID {
	return u.id
}
func (u *User) UserName() string {
	return u.userName
}
func (u *User) Password() string {
	return u.password
}

func (u *User) SetId(id uuid.UUID) {
	u.id = id
}
func (u *User) SetUserName(userName string) {
	u.userName = userName
}
func (u *User) SetPassword(password string) {
	u.password = password
}
