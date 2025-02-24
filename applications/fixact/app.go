package fixact

import (
	"database/sql"

	"fyne.io/fyne/v2"
	_ "github.com/mattn/go-sqlite3"
)

func NewApp(db *sql.DB) fyne.CanvasObject {
	createTableInDB(db)
	content := createInterfaceApp(db)
	return content
}

func OldApp(db *sql.DB) fyne.CanvasObject {
	content := createInterfaceApp(db)
	writeFieldApp(widgtsApp) // записываем последние значения
	return content
}
