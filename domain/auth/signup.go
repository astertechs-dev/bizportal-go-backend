package auth

import (
	"context"
	"github.com/astertechs-dev/bizportal-go-backend/domain/user"
)

type SignupRequest struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required, email"`
	Password string `form:"password" binding:"required"`
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignupUsecase interface {
	Create(c context.Context, user *user.User) error
	GetUserByEmail(c context.Context, email string) (*user.User, error)
	CreateAccessToken(user *user.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *user.User, secret string, expiry int) (refreshToken string, err error)
}
