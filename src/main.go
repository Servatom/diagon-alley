package main

import (
	deliver_rest "github.com/servatom/diagon-alley/src/delivery/http"
	"github.com/servatom/diagon-alley/src/internal/repository"
	"github.com/servatom/diagon-alley/src/internal/usecase"
	"github.com/servatom/diagon-alley/src/utils"
	postgres_database "github.com/servatom/diagon-alley/src/utils/database"
)

func main() {
	config := utils.NewConfig()
	db := postgres_database.NewGormPsqlClient(config)
	
	repositories := repository.InitRepositories(config, db)
	usecases := usecase.InitUsecases(config, repositories)

	deliver_rest.RestDeliver(config, usecases)
}