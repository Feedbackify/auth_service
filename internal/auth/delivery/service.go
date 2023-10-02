package delivery

import (
	"context"
	"fmt"
	"github.com/Feedbackify/auth_service/internal/auth/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"sync"

	auth_v1 "github.com/Feedbackify/auth_service/proto/auth/v1"
)

type AuthHandler struct {
	auth_v1.UnimplementedAuthServiceServer
	mu  *sync.RWMutex
	auc domain.AuthUseCase
}

func NewAuthHandler(gserver *grpc.Server, useCase domain.AuthUseCase) {

	authServer := &AuthHandler{
		mu:  &sync.RWMutex{},
		auc: useCase,
	}

	auth_v1.RegisterAuthServiceServer(gserver, authServer)
	reflection.Register(gserver)
}

func (a *AuthHandler) Login(ctx context.Context, request *auth_v1.LoginRequest) (*auth_v1.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthHandler) ResetPassword(ctx context.Context, request *auth_v1.ResetPasswordRequest) (*auth_v1.ResetPasswordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthHandler) ChangePassword(ctx context.Context, request *auth_v1.ChangePasswordRequest) (*auth_v1.ChangePasswordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthHandler) RefreshToken(ctx context.Context, request *auth_v1.RefreshTokenRequest) (*auth_v1.RefreshTokenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthHandler) Logout(ctx context.Context, request *auth_v1.LogoutRequest) (*auth_v1.LogoutResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthHandler) VerifyToken(ctx context.Context, request *auth_v1.VerifyTokenRequest) (*auth_v1.VerifyTokenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthHandler) Register(ctx context.Context, req *auth_v1.RegisterRequest) (*auth_v1.RegisterResponse, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	fmt.Println(ctx)
	fmt.Println(req)
	register := a.toEntity(req)
	token, _ := a.auc.Register(ctx, register)
	return &auth_v1.RegisterResponse{Code: "10", Status: "10", Data: &auth_v1.Data{AccessToken: token.AccessToken, RefreshToken: token.RefreshToken}}, nil
}

func (a *AuthHandler) toEntity(req *auth_v1.RegisterRequest) domain.RegisterRequest {
	return domain.RegisterRequest{Email: req.Email, Password: req.Password}
}
