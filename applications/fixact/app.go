package fixact

import (
	"database/sql"

	"fyne.io/fyne/v2"
	_ "github.com/mattn/go-sqlite3"
)

// структура описывающая данные приложения FIXACT
// НЕОБХОДИМА ДЛЯ ТОГО, ЧТОБЫ КАЖДЫЙ РАЗ НЕ ОБНОВЛЯТЬ ДАННЫЕ ПРИЛОЖЕНИЯ ПРИ ПЕРЕКЛЮЧЕНИИ МЕЖДУ МИКРОПРИЛОЖЕНИЯМИ
type FixActData struct {
	activityType string
	startTime    string
	endTime      string
	totalTime    string
	comment      string
}

// AppFixActData - AFD
var AFixD FixActData

// ↑↑↑↑
// ТЕПЕРЬ ОСТАЛОСЬ ПРИДУМАТЬ КОГДА МЫ СОЗДАЕМ ОБЪЕКТ ДАННОЙ СТРУКТУРЫ

func NewApp(db *sql.DB) fyne.CanvasObject {

	createTableInDB(db)
	content := createInterfaceApp(db)

	return content

}

// БУДЕТ
func updateStateFixAct() {
	AFixD = FixActData{
		activityType: "",
		startTime:    "",
		endTime:      "",
		totalTime:    "",
		comment:      "",
	}
}
