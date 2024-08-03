package services

import (
	"database/sql"
	"ingesoft/backend/internal/models"
	"ingesoft/backend/internal/repository"
)

var db *sql.DB

func InitDatabase(database *sql.DB) {
	db = database
}

func FetchData() ([]models.User, error) {
	return repository.GetData(db)
}
