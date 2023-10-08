package domain

import "context"

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type RegisterRequest struct {
	Email    string
	Password string
}

type AuthUseCase interface {
	Login()
	Register(context.Context, RegisterRequest) (Tokens, error)
}

type AuthRepository interface {
	StoreToken()
	RetrieveToken()
	DeleteToken()
}
