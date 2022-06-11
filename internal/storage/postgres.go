package storage

import (
	"github.com/laeis/test-graphQL/internal/storage/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewStorage() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=1234567 dbname=test-grapthql port=5430 sslmode=disable TimeZone=UTC"
	gorm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = gorm.AutoMigrate(&model.User{}, &model.ToDo{})
	if err != nil {
		return nil, err
	}

	return gorm, nil
}
