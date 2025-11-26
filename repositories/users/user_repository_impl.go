package repositories

import (
	"context"
	"tokogue-api/models/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	err := repository.DB.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user *domain.User
	err := repository.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error
	
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repository *UserRepositoryImpl) FindByID(ctx context.Context, userID string) (*domain.User, error) {
	var user *domain.User
	err := repository.DB.WithContext(ctx).Where("id = ?", userID).Take(&user).Error
	
	if err != nil {
		return nil, err
	}
	return user, nil
}