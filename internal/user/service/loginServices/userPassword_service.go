package loginservices

import "ch-gateway/internal/user/domain"

type UserPasswordLoginService struct {
	Repository domain.UserRepository
	SigningKey string
}

func (u UserPasswordLoginService) Authenticate(userName, password string) (domain.AuthResponse, error) {
	user, err := u.Repository.FindUserByUserName(userName)
	emptyAuth := domain.AuthResponse{}
	if err != nil {
		return emptyAuth, err
	}
	err = user.CheckPassword(password)
	if err != nil {
		return emptyAuth, err
	}
	token, err := GenerateToken(user.Id, u.SigningKey)
	if err != nil {
		return emptyAuth, err
	}
	response := domain.AuthResponse{UserID: user.Id, Token: token}
	return response, nil
}
