package repository_auth

import (
	"context"

	domain_auth "github.com/servatom/diagon-alley/src/internal/domain/auth"
	base_repository "github.com/servatom/diagon-alley/src/internal/repository/base"
	"github.com/servatom/diagon-alley/src/utils"
	"gorm.io/gorm"
)

type UserRepository struct {
	Email        string `json:"email" gorm:"type:varchar(100);not null;unique"`
	Password     string `json:"password" gorm:"type:varchar(100);not null"`
	IsAdmin      bool   `json:"is_admin" gorm:"default:false"`
	base_repository.BaseRepository
}

type AuthRepositoryImplementation struct {
	db     *gorm.DB
	config *utils.Config
}

func (UserRepository) TableName() string {
	return "user_table"
}

func (a *AuthRepositoryImplementation) CreateUser(
	newUser domain_auth.User,
) (*domain_auth.UserWithID, error) {
	newUserModel := NewUserRepository(&newUser)
	err := a.db.Create(&newUserModel).Error
	if err != nil {
		return nil, err
	}
	return newUserModel.toDomainUser(), nil
}

func (a *AuthRepositoryImplementation) GetUserByEmail(
	ctx context.Context,
	email string,
) (*domain_auth.UserWithID, error) {
	user := &UserRepository{}
	err := a.db.Where("email = ?", email).First(user).Error
	if err != nil {
		return nil, err
	}
	return user.toDomainUser(), nil
}

func (a *AuthRepositoryImplementation) UpdateUserToAdminRole(
	ctx context.Context,
	email string,
	isAdmin bool,
) (*domain_auth.UserWithID, error) {
	user := &UserRepository{}
	err := a.db.Where("email = ?", email).First(user).Error
	if err != nil {
		return nil, err
	}
	user.IsAdmin = isAdmin
	err = a.db.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user.toDomainUser(), nil
}

func (a *AuthRepositoryImplementation) GetAllUsers(
	ctx context.Context,
) ([]*domain_auth.UserWithID, error) {
	var users []*UserRepository
	err := a.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	var domainUsers []*domain_auth.UserWithID
	for _, user := range users {
		domainUsers = append(domainUsers, user.toDomainUser())
	}
	return domainUsers, nil
}

func NewAuthRepositoryImplementation(
	db *gorm.DB,
	config *utils.Config,
) *AuthRepositoryImplementation {
	err := db.AutoMigrate(&UserRepository{})
	if err != nil {
		panic(err)
	}
	return &AuthRepositoryImplementation{
		db:     db,
		config: config,
	}
}
