package loginservices

import "ch-gateway/internal/user/domain"

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

func (u UserPasswordLoginService) Authenticate(userName, password string) (domain.AuthResponse, error) {
	user, err := u.repository.FindUserByUserName(userName)
	emptyAuth := domain.AuthResponse{}
	if err != nil {
		return emptyAuth, err
	}
	err = user.CheckPassword(password)
	if err != nil {
		return emptyAuth, err
	}
	token, err := GenerateToken(user.Id, u.signingKey)
	if err != nil {
		return emptyAuth, err
	}
	response := domain.AuthResponse{UserID: user.Id, Token: token}
	return response, nil
}
