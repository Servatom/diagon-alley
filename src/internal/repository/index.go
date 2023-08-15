package repository

import (
	domain_auth "github.com/servatom/diagon-alley/src/internal/domain/auth"
	domain_product "github.com/servatom/diagon-alley/src/internal/domain/product"
	repository_auth "github.com/servatom/diagon-alley/src/internal/repository/auth"
	repository_product "github.com/servatom/diagon-alley/src/internal/repository/product"
	"github.com/servatom/diagon-alley/src/utils"
	"gorm.io/gorm"
)

type Repositories struct {
	Auth domain_auth.Repository
	Product domain_product.Repository
}

func InitRepositories(
	config *utils.Config,
	db *gorm.DB,
) *Repositories {
	return &Repositories{
		Auth: repository_auth.NewAuthRepositoryImplementation(db, config),
		Product: repository_product.NewProductRepositoryImplementation(db, config),
	}
}