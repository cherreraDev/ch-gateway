package httprequests

import "github.com/google/uuid"

type CreateUserRequest struct {
	Id       uuid.UUID `json:"id"`
	UserName string    `json:"user_name"`
	Password string    `json:"password"`
}

type UpdateUserRequest struct {
	Id       uuid.UUID `json:"id"`
	UserName string    `json:"user_name"`
	Password string    `json:"password"`
}

type DeleteUserRequest struct {
	Id uuid.UUID `json:"id"`
}
