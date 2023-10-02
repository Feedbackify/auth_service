package application

import (
	"context"
	"github.com/Feedbackify/auth_service/internal/auth/domain"
)

type AuthUseCase struct {
	userRepo string
	authRepo string
}

func NewAuthUseCase() domain.AuthUseCase {
	return &AuthUseCase{authRepo: "", userRepo: "s"}
}

func (a *AuthUseCase) Login() {

}

func (a *AuthUseCase) Register(ctx context.Context, req domain.RegisterRequest) (domain.Tokens, error) {
	return domain.Tokens{AccessToken: "vvvv", RefreshToken: "11111111"}, nil
}

func (a *AuthUseCase) ChangePassword() {

}
