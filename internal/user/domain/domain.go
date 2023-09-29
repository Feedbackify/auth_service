package domain

import (
	"github.com/google/uuid"
	"time"
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

func (u *User) hashPassword(password string) string {
	return password
}

func (u *User) New(userDto UserDto) *User {
	now := time.Now()
	return &User{ID: uuid.New(), Email: userDto.Email, Password: u.hashPassword(userDto.Password), Profile: userDto.Profile, Created: now, Updated: now}
}
