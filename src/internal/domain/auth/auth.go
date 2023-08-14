package domain_auth

import (
	"context"
)

type User struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	IsAdmin 	bool   `json:"is_admin"`
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
	ID           int64  `json:"id"`
	Email        string `json:"email"`
	IsAdmin      bool   `json:"is_admin"`
}

type Usecase interface {
	Login(ctx context.Context, username string, password string) (token *string, err error)
	VerifyToken(ctx context.Context, token string) (*UserWithID, error)
	UpdateUserToAdminRole(ctx context.Context, email string, isAdmin bool) (bool, error)
	CreateUser(ctx context.Context, email string, password string) (bool, error)
	GetAllUsers(ctx context.Context) ([]*UserDetails, error)
}

type Repository interface {
	GetUserByEmail(ctx context.Context, email string) (*UserWithID, error)
	CreateUser(newUser User) (*UserWithID, error)
	UpdateUserToAdminRole(ctx context.Context, email string, isAdmin bool) (*UserWithID, error)
	GetAllUsers(ctx context.Context) ([]*UserWithID, error)
}
