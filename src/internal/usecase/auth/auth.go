package usecase_auth

import (
	"context"
	"errors"
	"net/mail"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	domain_auth "github.com/servatom/diagon-alley/src/internal/domain/auth"
	"github.com/servatom/diagon-alley/src/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecaseImplementation struct {
	contextTimeout time.Duration
	config         *utils.Config
	authRepository domain_auth.Repository
}

type CustomClaims struct {
	domain_auth.UserWithID
	jwt.StandardClaims
}

func (a *AuthUsecaseImplementation) isFirstUser(
	ctx context.Context,
) (bool, error) {
	users, err := a.authRepository.GetAllUsers(ctx)
	if err != nil {
		return false, err
	}
	return len(users) == 0, nil
}

func (a *AuthUsecaseImplementation) CreateUser(
	ctx context.Context,
	email string,
	password string,
	gender string,
	age int,
) (bool, error) {
	// validate email
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false, err
	}
	userCheck, _ := a.authRepository.GetUserByEmail(ctx, strings.ToLower(email))
	if userCheck != nil {
		return false, errors.New("\nuser already exists")
	}
	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}

	isAdmin := false

	isFirstUser, err := a.isFirstUser(ctx)
	if err != nil {
		return false, err
	}
	if isFirstUser {
		isAdmin = true
	}
	user := domain_auth.User{
		Email:        strings.ToLower(email),
		Password:     string(hashedPassword),
		IsAdmin:      isAdmin,
		Gender: gender,
		Age: age,
	}
	_, err = a.authRepository.CreateUser(user)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *AuthUsecaseImplementation) Login(
	ctx context.Context,
	email string,
	password string,
) (*string, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()
	// get user
	user, err := a.authRepository.GetUserByEmail(ctx, strings.ToLower(email))
	if err != nil {
		return nil, err
	}
	// check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	// generate token
	claims := CustomClaims{
		*user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			Issuer:    "diagon-alley",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(a.config.SecretKey))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func (a *AuthUsecaseImplementation) VerifyToken(
	ctx context.Context,
	token string,
) (*domain_auth.UserWithID, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()
	// parse token
	parsedToken, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.config.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	// check if token is valid
	claims, ok := parsedToken.Claims.(*CustomClaims)
	if !ok || !parsedToken.Valid {
		return nil, err
	}
	// get user
	user, err := a.authRepository.GetUserByEmail(ctx, claims.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a *AuthUsecaseImplementation) UpdateUserToAdminRole(
	ctx context.Context,
	email string,
	isAdmin bool,
) (bool, error) {
	_, err := a.authRepository.UpdateUserToAdminRole(ctx, strings.ToLower(email), isAdmin)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *AuthUsecaseImplementation) GetAllUsers(
	ctx context.Context,
) ([]*domain_auth.UserDetails, error) {
	users, err := a.authRepository.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	var userDetails []*domain_auth.UserDetails
	for _, user := range users {
		userDetails = append(userDetails, &domain_auth.UserDetails{
			ID:           user.ID,
			Email:        user.Email,
			IsAdmin:      user.IsAdmin,
			Gender:       user.Gender,
			Age:          user.Age,
		})
	}
	return userDetails, nil
}

func (a *AuthUsecaseImplementation) GetMe(
	ctx context.Context,
	user_id int64,
) (*domain_auth.UserProfile, error) {
	user, err := a.authRepository.GetUserById(ctx, user_id)
	if err != nil {
		return nil, err
	}
	var finalResponse domain_auth.UserProfile
	if user.Gender == "f" {
		finalResponse.Gender = "Female"
	}else if user.Gender == "m" {
		finalResponse.Gender = "Male"
	}else{
		finalResponse.Gender = "other"
	}
	finalResponse.Age = user.Age
	finalResponse.Name = user.Name
	return &finalResponse, nil
}

func NewAuthUsecaseImplementation(
	config *utils.Config,
	authRepository domain_auth.Repository,
) *AuthUsecaseImplementation {
	return &AuthUsecaseImplementation{
		config:         config,
		authRepository: authRepository,
	}
}
