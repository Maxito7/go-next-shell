package main

import (
	"database/sql"
	"ingesoft/backend/config"
	"ingesoft/backend/internal/server"
	"ingesoft/backend/internal/services"
	"log"
)

func main() {
	// Inicializamos la BD
	dbDriver := config.GetEnv("DB_DRIVER", "sqlite3")
	dbSource := config.GetEnv("DB_SOURCE", "../../db/cascaron.db")
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	services.InitDatabase(db)

	server.Start()
}
