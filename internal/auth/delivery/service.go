package delivery

import (
	"context"
	"fmt"
	auth_v1 "github.com/Feedbackify/auth_service/proto/auth/v1"
	"sync"
)

type Backend struct {
	auth_v1.UnimplementedAuthServiceServer
	mu *sync.RWMutex
}

func (b *Backend) Login(ctx context.Context, request *auth_v1.LoginRequest) (*auth_v1.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (b *Backend) ResetPassword(ctx context.Context, request *auth_v1.ResetPasswordRequest) (*auth_v1.ResetPasswordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (b *Backend) ChangePassword(ctx context.Context, request *auth_v1.ChangePasswordRequest) (*auth_v1.ChangePasswordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (b *Backend) RefreshToken(ctx context.Context, request *auth_v1.RefreshTokenRequest) (*auth_v1.RefreshTokenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (b *Backend) Logout(ctx context.Context, request *auth_v1.LogoutRequest) (*auth_v1.LogoutResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (b *Backend) VerifyToken(ctx context.Context, request *auth_v1.VerifyTokenRequest) (*auth_v1.VerifyTokenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (b *Backend) mustEmbedUnimplementedAuthServiceServer() {
	//TODO implement me
	panic("implement me")
}

func New() *Backend {
	return &Backend{
		mu: &sync.RWMutex{},
	}
}

func (b *Backend) Register(ctx context.Context, _ *auth_v1.RegisterRequest) (*auth_v1.RegisterResponse, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	fmt.Println("read it ")
	return &auth_v1.RegisterResponse{Code: "10", Status: "10", Data: &auth_v1.Data{AccessToken: "vvvv", RefreshToken: "asdasdas"}}, nil
}
