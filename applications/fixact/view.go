package fixact

import (
	"database/sql"
	"fmt"
	"mainjournal/applications/fixact/pos"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

func createInterfaceApp(db *sql.DB) (content *fyne.Container) {
	// Элементы интерфейса
	activityType := widget.NewSelect([]string{"Книга", "Код", "Видео"}, func(value string) { activityTypeEvent(value) })
	activityType.PlaceHolder = "Выбери активность"

	startTime := widget.NewEntry()
	startTime.SetText(getNow())

	endTime := widget.NewEntry()
	totalTime := widget.NewEntry()
	comment := widget.NewMultiLineEntry()
	addButton := widget.NewButton("Добавить активность", func() { addAct(widgtsApp, db) })

	// ДУБЛИРУЕМ СОЗДАННЫЕ ВИДЖЕТЫ В ГЛОБАЛЬНУЮ ПЕРЕМЕННУЮ
	widgtsApp = Widgets{
		activityType: activityType,
		startTime:    startTime,
		endTime:      endTime,
		totalTime:    totalTime,
		comment:      comment,
		addButton:    addButton,
	}

	btnSupp1 := widget.NewButton("🞴", func() { btnSupp1Event(widgtsApp) })
	btnSupp2 := widget.NewButton("🞴", func() { btnSupp2Event(widgtsApp) })
	btnSupp3 := widget.NewButton("🞴", func() { btnSupp3Event(widgtsApp) })

	globContainer := container.NewWithoutLayout()

	h1 := float32(10)
	h2 := float32(50)
	h3 := float32(90)
	h4 := float32(130)
	h5 := float32(170)
	h6 := float32(250)

	pos.AddRow(globContainer, WIDTH, h1, widget.NewLabel("Тип активности:"), activityType)
	pos.AddRow(globContainer, WIDTH, h2, widget.NewLabel("Время начала:"), startTime, btnSupp1)
	pos.AddRow(globContainer, WIDTH, h3, widget.NewLabel("Время окончания:"), endTime, btnSupp2)
	pos.AddRow(globContainer, WIDTH, h4, widget.NewLabel("Общее время:"), totalTime, btnSupp3)
	pos.AddRow(globContainer, WIDTH, h5, widget.NewLabel("Комментарий:"), comment)
	pos.AddRow(globContainer, WIDTH, h6, widget.NewLabel(""), addButton)

	return globContainer
}

// activityType EVENTS (CLCIK)
func activityTypeEvent(value string) {
	fmt.Println("Выбран ", value)
}

// ФУНКЦИИ ДЛЯ РАБОТЫ ВСПОМОГАТЕЛЬНЫХ КНОПОК
// ДЛЯ НИХ НУЖНО ПИСАТЬ ГЛОБАЛЬНЫЙ ОБЪЕКТ ВИДЖЕТОВ ПРИЛОЕЖНИЯ, ЧТОБЫ В ЭТИХ ФУНКЦИЯХ
// СЧИТЫВАТЬ ЗНАЧЕНИЯ С ЭТИХ ВИДЖЕТОВ И ПОМЕЩАТЬ ИХ В AFixD (объект состояний полей ввода ФИКС-АКТА)
func btnSupp1Event(widgtsApp Widgets) {
	widgtsApp.startTime.SetText(getNow())
	appFieldVal = readFieldApp(widgtsApp)
	fmt.Println(appFieldVal)

}

func btnSupp2Event(widgtsApp Widgets) {
	widgtsApp.endTime.SetText(getNow())
	appFieldVal = readFieldApp(widgtsApp)
	fmt.Println(appFieldVal)
}

func btnSupp3Event(widgtsApp Widgets) {
	widgtsApp.totalTime.SetText(getActTime(widgtsApp.endTime.Text, widgtsApp.startTime.Text))
	appFieldVal = readFieldApp(widgtsApp)
	fmt.Println(appFieldVal)
}
