package repository

import (
	"database/sql"
	"ingesoft/backend/internal/models"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func GetData(db *sql.DB) ([]models.User, error) {
	rows, err := db.Query("SELECT id, name, lastname, codpucp FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.LastName, &user.CodPucp); err != nil {
			log.Fatalf("Error getting user data from the DB: %v", err)
			return nil, err
		}
		data = append(data, user)
	}
	return data, nil
}
