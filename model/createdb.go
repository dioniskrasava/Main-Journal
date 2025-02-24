package dataApp

import (
	"database/sql"
	"log"
)

func CreateDB() *sql.DB {
	var err error
	var db *sql.DB

	// Если путь к базе данных не указан, используем путь по умолчанию
	db, err = sql.Open("sqlite3", "./activities.db")
	if err != nil {
		log.Fatal(err)
	}

	// Проверяем подключение к базе данных
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	return db
}
