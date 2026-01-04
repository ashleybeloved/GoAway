package repositories

import (
	"fmt"
	"goaway/internal/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func StartPostgreSQL() error {
	// Connecting to PostgreSQL server

	host := os.Getenv("POSTGRESQL_HOST")
	user := os.Getenv("POSTGRESQL_USER")
	password := os.Getenv("POSTGRESQL_PASSWORD")
	dbname := os.Getenv("POSTGRESQL_DB_NAME")
	port := os.Getenv("POSTGRESQL_PORT")
	sslmode := os.Getenv("POSTGRESQL_SSL_MODE")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v", host, user, password, dbname, port, sslmode)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Creating if exists two tables with structure models.User{} and models.Link{}

	err := db.AutoMigrate(&models.User{}, &models.Link{})
	if err != nil {
		return err
	}

	return nil
}

func CreateUser(login string, hashedPassword []byte) error {
	user := models.User{
		Login:    login,
		Password: string(hashedPassword),
	}

	return db.Create(&user).Error
}

func FindUserByLogin(login string) (*models.User, error) {
	var user models.User

	result := db.Where("login = ?", login).Find(&user)
	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &user, result.Error
}
