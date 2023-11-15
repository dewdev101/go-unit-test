package config

import (
	"fmt"

	"github.com/ThanaponBunchot/go-unit-test/migrations"
	"github.com/ThanaponBunchot/go-unit-test/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dsn := "host=localhost user=root password=root dbname=postgres port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("unable to conect database")
	}
	migrations.AutoMigrate(db, &model.Product{})
	fmt.Println("database connnected!")
	return db
}

func ConnectDatabaseTest() *gorm.DB {
	dsn := "host=localhost user=root password=root dbname=postgres port=5434 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("unable to conect database")
	}
	migrations.AutoMigrate(db, &model.Product{})
	fmt.Println("database test connnected!")
	return db
}
