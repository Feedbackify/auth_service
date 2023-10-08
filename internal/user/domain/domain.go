package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserProfile struct {
	Name   string
	Family string
}

type User struct {
	ID       uuid.UUID
	Email    string
	Password string
	Profile  UserProfile
	Created  time.Time
	Updated  time.Time
}

type UserDto struct {
	Email    string
	Password string
	Profile  UserProfile
}

type UserUseCase interface {
	Create()
	List()
	GetOne()
}

type UserRepository interface {
	Store()
	FindAll()
	FindById()
}
