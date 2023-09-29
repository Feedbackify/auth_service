package application

import "github.com/Feedbackify/auth_service/internal/auth/domain"

type AuthUseCase struct {
	userRepo string
	authRepo string
}

func NewAuthUseCase() domain.AuthUseCase {
	return &AuthUseCase{authRepo: "", userRepo: "s"}
}

func (a AuthUseCase) Login() {

}

func (a AuthUseCase) Register() {

}

func (a AuthUseCase) ChangePassword() {

}
