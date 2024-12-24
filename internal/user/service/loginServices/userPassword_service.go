package loginservices

import (
	"ch-gateway/internal/user/domain"
	"errors"
)

type UserPasswordLoginService struct {
	repository domain.UserRepository
	signingKey string
}

func NewUserPasswordLoginService(repository domain.UserRepository, signingKey string) UserPasswordLoginService {
	return UserPasswordLoginService{
		repository: repository,
		signingKey: signingKey,
	}
}

func (u UserPasswordLoginService) Authenticate(credentials map[string]string) (domain.AuthResponse, error) {
	userName, userOk := credentials["username"]
	password, passOk := credentials["password"]

	if !userOk || !passOk {
		return domain.AuthResponse{}, errors.New("faltan credenciales")
	}
	user, err := u.repository.FindUserByUserName(userName)
	emptyAuth := domain.AuthResponse{}
	if err != nil {
		return emptyAuth, err
	}
	err = user.CheckPassword(password)
	if err != nil {
		return emptyAuth, err
	}
	token, err := GenerateToken(user.Id(), u.signingKey)
	if err != nil {
		return emptyAuth, err
	}
	response := domain.AuthResponse{UserID: user.Id(), Token: token}
	return response, nil
}
