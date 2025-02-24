package fixact

import (
	"database/sql"

	"fyne.io/fyne/v2"
	_ "github.com/mattn/go-sqlite3"
)

// ВЕРСИЯ  V 1.1.001 - ДОБАВИТЬ ФИЧУ (ВЫБОР МЕСТА СОХРАНИЕНИЯ БАЗЫ ДАННЫХ)

func NewApp(db *sql.DB) fyne.CanvasObject {

	createTableInDB(db)
	content := createInterfaceApp(db)

	return content

}
