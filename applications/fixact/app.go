package fixact

import (
	"database/sql"

	"fyne.io/fyne/v2"
	_ "github.com/mattn/go-sqlite3"
)

// копия главной базы данных
// чтобы юзать её внутри
// а то тут некоторые ф-ии без неё работать глядикось не хотят!
var maindb *sql.DB // вызывается в drawFixActOnMainWindow()

func NewApp(db *sql.DB) fyne.CanvasObject {
	maindb = db
	createTableInDB(db)
	content := createInterfaceApp(db)
	return content
}

func OldApp(db *sql.DB) fyne.CanvasObject {
	maindb = db
	content := createOldInterfaceApp(db)
	writeFieldApp(widgtsApp) // записываем последние значения
	return content
}
