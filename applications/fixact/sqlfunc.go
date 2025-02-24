package fixact

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func createTableInDB(db *sql.DB) {

	// Создание таблицы, если она не существует
	createTableSQL := `CREATE TABLE IF NOT EXISTS activities (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	type TEXT,
	start_time TEXT,
	end_time TEXT,
	total_time TEXT,
	comment TEXT
);`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

func addAct(w Widgets, db *sql.DB) {
	activity := readFieldApp(w)

	// Вставка данных в базу данных
	insertSQL := `INSERT INTO activities (type, start_time, end_time, total_time, comment) VALUES (?, ?, ?, ?, ?)`
	_, err := db.Exec(insertSQL, activity.Type, activity.StartTime, activity.EndTime, activity.TotalTime, activity.Comment)
	if err != nil {
		log.Fatal(err)
	}
	cleanFielApp(w)
	fmt.Println("Активность добавлена!")
}
