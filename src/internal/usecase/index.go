package usecase

import (
	domain_auth "github.com/servatom/diagon-alley/src/internal/domain/auth"
	"github.com/servatom/diagon-alley/src/internal/repository"
	"github.com/servatom/diagon-alley/src/utils"
	usecase_auth "github.com/servatom/diagon-alley/src/internal/usecase/auth"
)

type Usecases struct {
	Auth domain_auth.Usecase
}

func InitUsecases(
	config *utils.Config,
	repositories *repository.Repositories,
) *Usecases {
	return &Usecases{
		Auth: usecase_auth.NewAuthUsecaseImplementation(config, repositories.Auth),
	}
}