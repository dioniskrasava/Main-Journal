package fixact

import (
	"fyne.io/fyne/v2/widget"
)

const (
	WIDTH  float32 = 425
	HEIGHT float32 = 350
)

// структура для сбора данных перед отправкой в БАЗУ ДАННЫХ
type Activity struct {
	Type      string
	StartTime string
	EndTime   string
	TotalTime string
	Comment   string
}

// структура виджетов микроприложения
type Widgets struct {
	activityType *widget.Select
	startTime    *widget.Entry
	endTime      *widget.Entry
	totalTime    *widget.Entry
	comment      *widget.Entry
	addButton    *widget.Button
}

var widgtsApp Widgets = Widgets{} // виджеты FixAct
var appFieldVal = Activity{}

// ПЕРЕМЕННАЯ БАЗЫ ДАННЫХ ГЛОБАЛЬНАЯ НА УРОВНЕ ПАКЕТА - ВОЗМОЖНО ЭТО НЕ ПРАВИЛЬНО
// НО К СОЖАЛЕНИЮ ПОКА ЭТО ОДИН ИЗ УДОБНЫХ ВАРИАНТОВ ЭТОГО КОДА ИНАЧЕ НУЖНО БЫЛО БЫ
// ОБДУМЫВАТЬ БОЛЕЕ ТЩАТЕЛЬНО АРХИТЕКТУРУ ПРИЛОЖЕНИЯ
// Проблема в том, что при старте приложения мы подключались (создавали) бд рядом с приложением
// и в дальнейшем при поптыке подключить другую бд Я сталкивался с проблемой того, что запись по прежнему
// производилась в старый файл бд, а не выбранный мною
