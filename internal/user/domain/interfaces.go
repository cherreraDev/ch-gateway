package domain

import "github.com/google/uuid"

type LoginService interface {
	Authenticate(credentials ...string) (AuthResponse, error)
}

type AuthResponse struct {
	UserID uuid.UUID
	Token  string
}

type UserRepository interface {
	FindUserById(userId uuid.UUID) (User, error)
	FindUserByUserName(userName string) (User, error)
	SaveUser(User) error
	UpdateUser(User) error
	DeleteUser(userId uuid.UUID) error
}
