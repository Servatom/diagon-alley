package main

import (
	"context"
	"fmt"
	"net/mail"
	"syscall"

	repository_auth "github.com/servatom/diagon-alley/src/internal/repository/auth"
	usecase_auth "github.com/servatom/diagon-alley/src/internal/usecase/auth"
	utils "github.com/servatom/diagon-alley/src/utils"
	database_postgres "github.com/servatom/diagon-alley/src/utils/database"
	"golang.org/x/term"
)

func main() {
	ctx := context.Background()
	config := utils.NewConfig()
	db := database_postgres.NewGormPsqlClient(config)
	authRepo := repository_auth.NewAuthRepositoryImplementation(db, config)
	authUsecase := usecase_auth.NewAuthUsecaseImplementation(config, authRepo)
	var email string
	fmt.Print("Enter your email: ")
	fmt.Scanln(&email)
	_, err := mail.ParseAddress(email)
	if err != nil {
		fmt.Println("Invalid email")
		return
	}

	var password string
	fmt.Print("Enter your password: ")
	bytePasswd, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		fmt.Println("Error while reading password")
		return
	}

	password = string(bytePasswd)
	if len(password) < 8 {
		fmt.Println("Password should be atleast 8 characters long")
		return
	}

	// validate confirm password
	var confirmPassword string
	fmt.Print("\nConfirm your password: ")
	bytePasswd, err = term.ReadPassword(syscall.Stdin)
	if err != nil {
		fmt.Println("Error while reading password")
		return
	}
	confirmPassword = string(bytePasswd)
	if confirmPassword != password {
		fmt.Println("Passwords do not match")
		return
	}

	_, err = authUsecase.CreateUser(ctx, email, password)
	if err != nil {
		return
	}
	fmt.Println("\nUser created successfully")
}
