package services

import (
	"context"
	"tokogue-api/models/web"
)

type AuthService interface {
	Register(ctx context.Context, request web.UserCreateRequest) (web.UserResponse, error)
	Login(ctx context.Context, request web.UserLoginRequest) (web.UserLoginResponse, error)
}