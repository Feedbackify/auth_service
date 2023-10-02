package repository

import "github.com/Feedbackify/auth_service/internal/auth/domain"

type AuthRepository struct {
	db   string
	name string
}

func NewAuthRepository() domain.AuthRepository {
	return &AuthRepository{db: "db connection", name: "test"}
}

func (a *AuthRepository) StoreToken() {

}

func (a *AuthRepository) RetrieveToken() {

}

func (a *AuthRepository) DeleteToken() {

}
