package fixact

import (
	"fyne.io/fyne/v2"
	_ "github.com/mattn/go-sqlite3"
)

// ВЕРСИЯ  V 1.1.001 - ДОБАВИТЬ ФИЧУ (ВЫБОР МЕСТА СОХРАНИЕНИЯ БАЗЫ ДАННЫХ)

func NewApp() fyne.CanvasObject {

	createDB()
	defer db.Close()

	createTableInDB()
	content := createInterfaceApp()

	return content

}
