package entity

import (
	"github.com/0x16F/cloud-common/pkg/generator"
)

const (
	SaltLength = 10
)

type User struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"-"`
	Salt     string `json:"-"`
}

type UserCreateDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetUsersParams struct {
	Limit    int
	LastID   uint64
	Username string
	Email    string
}

type UserData struct {
	Login string
	Role  string
}

func NewUser(dto UserCreateDTO) User {
	salt := generator.NewString(SaltLength)
	password := generator.NewHash(dto.Password, salt)

	return User{
		Email:    dto.Email,
		Username: dto.Username,
		Password: password,
		Salt:     salt,
	}
}

func (u User) ValidatePassword(password string) bool {
	return generator.NewHash(password, u.Salt) == u.Password
}
