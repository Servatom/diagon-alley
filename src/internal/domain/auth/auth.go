package domain_auth

import (
	"context"
)

type User struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	IsAdmin 	 bool   `json:"is_admin"`
	Gender       string `json:"gender"`
	Age 		int    `json:"age"`
}

type EmailLoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserWithID struct {
	ID   int64 `json:"id"`
	User `json:",inline"`
}
type UserDetails struct {
	Name         string `json:"name"`
	ID           int64  `json:"id"`
	Email        string `json:"email"`
	IsAdmin      bool   `json:"is_admin"`
	Gender       string `json:"gender"`
	Age          int    `json:"age"`
}

type UserProfile struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

type Usecase interface {
	Login(ctx context.Context, username string, password string) (token *string, err error)
	VerifyToken(ctx context.Context, token string) (*UserWithID, error)
	UpdateUserToAdminRole(ctx context.Context, email string, isAdmin bool) (bool, error)
	CreateUser(ctx context.Context, email string, password string, gender string, age int) (bool, error)
	GetAllUsers(ctx context.Context) ([]*UserDetails, error)
	GetMe(ctx context.Context, user int64) (*UserProfile, error)
}

type Repository interface {
	GetUserByEmail(ctx context.Context, email string) (*UserWithID, error)
	CreateUser(newUser User) (*UserWithID, error)
	UpdateUserToAdminRole(ctx context.Context, email string, isAdmin bool) (*UserWithID, error)
	GetAllUsers(ctx context.Context) ([]*UserWithID, error)
	GetUserById(ctx context.Context, id int64) (*UserWithID, error)
}
