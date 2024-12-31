package domain

import "github.com/google/uuid"

type LoginService interface {
	Authenticate(credentials map[string]string) (AuthResponse, error)
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

type UserService interface {
	GetUserById(userId uuid.UUID) (User, error)
	GetUserByUserName(userName string) (User, error)
	CreateUser(id uuid.UUID, userName, password string) error
	UpdateUser(userName, password string) error
	DeleteUser(userId uuid.UUID) error
}
