package fyneview

import (
	"database/sql"
	"mainjournal/applications/fixact"
	dataApp "mainjournal/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

// глобальные переменные
type globalVariables struct {
	a           fyne.App
	w           fyne.Window
	splitGlobal *container.Split
	db          *sql.DB
}

// глобальные флаги
type globalFlags struct {
	beginStopwatch bool
}

var g_v globalVariables
var g_f globalFlags

func App() {
	g_v.a = app.NewWithID("Main Journal")
	g_v.w = g_v.a.NewWindow("Main Journal v 0.0.1")
	g_v.w.Resize(fyne.NewSize(480, 300)) // размер окна
	//w.SetFixedSize(true)
	g_v.w.CenterOnScreen()

	// СОЗДАЕМ ОБЩУЮ БАЗУ ДАННЫХ ПРИЛОЖЕНИЯ
	g_v.db = dataApp.CreateDB()
	defer g_v.db.Close() // не забудем закрыть вместе с приложением

	// Создаем два виджета, которые будут размещены в разделяемых панелях

	// БАГ - приложение ФИКСАКТ инициализируется МНОГО РАЗ!!! Нужно написать флаг
	right := fixact.NewApp(g_v.db)
	cont := createSideBar()

	// Создаем контейнер с разделяемыми панелями
	g_v.splitGlobal = container.NewHSplit(cont, right)
	g_v.splitGlobal.SetOffset(0.01) // Устанавливаем разделитель в середину

	// Устанавливаем контейнер в окно
	g_v.w.SetContent(g_v.splitGlobal)

	g_v.w.ShowAndRun()
}
