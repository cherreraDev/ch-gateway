package repositories

import (
	"ch-gateway/internal/user/domain"
	"ch-gateway/internal/user/platform/storage/model"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) GormUserRepository {
	return GormUserRepository{db: db}
}

func (r GormUserRepository) FindUserById(userId uuid.UUID) (domain.User, error) {
	var userModel model.UserModel
	if err := r.db.First(&userModel, "id = ?", userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}

	return model.MapToDomain(userModel), nil
}

func (r GormUserRepository) FindUserByUserName(userName string) (domain.User, error) {
	var userModel model.UserModel
	if err := r.db.First(&userModel, "user_name = ?", userName).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}

	return model.MapToDomain(userModel), nil
}

func (r GormUserRepository) SaveUser(user domain.User) error {
	userModel := model.MapToModel(user)
	return r.db.Create(&userModel).Error
}

func (r GormUserRepository) UpdateUser(user domain.User) error {
	userModel := model.MapToModel(user)
	return r.db.Save(&userModel).Error
}

func (r GormUserRepository) DeleteUser(userId uuid.UUID) error {
	return r.db.Delete(&model.UserModel{}, "id = ?", userId).Error
}
