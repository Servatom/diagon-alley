package repository

import (
	domain_auth "github.com/servatom/diagon-alley/src/internal/domain/auth"
	repository_auth "github.com/servatom/diagon-alley/src/internal/repository/auth"
	"github.com/servatom/diagon-alley/src/utils"
	"gorm.io/gorm"
)

type Repositories struct {
	Auth domain_auth.Repository
}

func InitRepositories(
	config *utils.Config,
	db *gorm.DB,
) *Repositories {
	return &Repositories{
		Auth: repository_auth.NewAuthRepositoryImplementation(db, config),
	}
}