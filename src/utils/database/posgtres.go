package postgres_database

import (
	"github.com/servatom/diagon-alley/src/utils"
	postgres "gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
)

func NewGormPsqlClient(config *utils.Config) *gorm.DB {
	dsn := utils.GetPostgresConnectionString(config)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
