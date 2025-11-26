package services

import (
	"context"
	"errors"
	"tokogue-api/config"
	"tokogue-api/exceptions"
	"tokogue-api/models/domain"
	"tokogue-api/models/web"
	repositories "tokogue-api/repositories/users"
	"tokogue-api/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	UserRepository repositories.UserRepository
	DB             *gorm.DB
	Config         *config.Config
}

func NewAuthServiceImpl(userRepository repositories.UserRepository, db *gorm.DB, config *config.Config) *AuthServiceImpl {
	return &AuthServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		Config:         config,
	}
}


func (service *AuthServiceImpl) Register(ctx context.Context, request web.UserCreateRequest) (web.UserResponse, error) {
	// 1. Check if email already exists
	existingUser, err := service.UserRepository.FindByEmail(ctx, request.Email)
	if err == nil && existingUser != nil {
		return web.UserResponse{}, exceptions.NewDuplicateError("Email sudah terdaftar")
	}
	// Ignore error if user not found (which is what we want)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return web.UserResponse{}, err
	}

	// 2. hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return web.UserResponse{}, err
	}

	user := &domain.User{
		FullName: request.FullName,
		Email:    request.Email,
		Password: string(hashedPassword),
		Role:     request.Role,
	}

	// 3. save user to database
	err = service.DB.Transaction(func(tx *gorm.DB) error {
		savedUser, err := service.UserRepository.Create(ctx, user)
		if err != nil {
			return err
		}
		if savedUser.ID == "" {
			return gorm.ErrInvalidData
		}
		return nil
	})

	if err != nil {
		return web.UserResponse{}, err
	}

	return web.UserResponse{
		FullName: user.FullName,
		Email:    user.Email,
		Role:     user.Role,
	}, nil
}

func (service *AuthServiceImpl) Login(ctx context.Context, request web.UserLoginRequest) (web.UserLoginResponse, error) {
	
	// 1. Cari User dari email
	user, err := service.UserRepository.FindByEmail(ctx, request.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return web.UserLoginResponse{}, exceptions.NewUnauthorizedError("invalid email")
		}
		return web.UserLoginResponse{}, err
	}

	// 2. Validasi Password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil {
		return web.UserLoginResponse{}, exceptions.NewUnauthorizedError("invalid password")
	}

	// 3. Generate Token JWT 
	token, errToken := utils.GenerateToken(user.ID, user.Role, service.Config.JWT.Secret)
	
	if errToken != nil {
		return web.UserLoginResponse{}, errToken
	}

	return web.UserLoginResponse{
		FullName: user.FullName,
		Email:    user.Email,
		Role:     user.Role,
		Token:    token,
	}, nil

}