package model

import (
	"ch-gateway/internal/user/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	ID       uuid.UUID `gorm:"type:char(36);primaryKey" `
	UserName string    `gorm:"type:varchar(100);not null;unique"`
	Password string    `gorm:"type:varchar(255);not null" `
}

func (u *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}
func MapToDomain(model UserModel) domain.User {
	return domain.NewUserBuilder().
		WithId(model.ID).
		WithUserName(model.UserName).
		WithPassword(model.Password).
		Build()
}

func MapToModel(user domain.User) UserModel {
	return UserModel{
		ID:       user.Id(),
		UserName: user.UserName(),
		Password: user.Password(),
	}
}
