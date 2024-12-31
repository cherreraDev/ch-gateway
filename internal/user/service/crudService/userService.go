package crudservice

import (
	"ch-gateway/internal/user/domain"
	"errors"

	"github.com/google/uuid"
)

type UserService struct {
	repository domain.UserRepository
}

func NewUserService(repository domain.UserRepository) UserService {
	return UserService{repository: repository}
}

func (us UserService) GetUserById(userId uuid.UUID) (domain.User, error) {
	return us.repository.FindUserById(userId)
}

func (us UserService) GetUserByUserName(userName string) (domain.User, error) {
	return us.repository.FindUserByUserName(userName)
}

func (us UserService) CreateUser(id uuid.UUID, userName, password string) error {
	user := domain.NewUserBuilder().WithId(id).WithUserName(userName).WithPassword(password).Build()
	err := user.EncryptPassword()
	if err != nil {
		return err
	}
	err = us.repository.SaveUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (us UserService) UpdateUser(id uuid.UUID, userName, password string) error {
	userBuilder := domain.NewUserBuilder()
	if userName != "" {
		userBuilder.WithUserName(userName)
	}
	if password != "" {
		userBuilder.WithPassword(password)
	}
	if id == uuid.Nil {
		return errors.New("invalid ID: ID cannot be nil")
	}
	userBuilder.WithId(id)
	user := userBuilder.Build()
	err := user.EncryptPassword()
	if err != nil {
		return err
	}
	err = us.repository.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (us UserService) DeleteUser(userId uuid.UUID) error {
	return us.repository.DeleteUser(userId)
}
