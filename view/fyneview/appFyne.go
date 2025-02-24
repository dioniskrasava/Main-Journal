package fyneview

import (
	"database/sql"
	"mainjournal/applications/fixact"
	dataApp "mainjournal/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

// структура глобальных переменных
type globalVariables struct {
	a           fyne.App
	w           fyne.Window
	splitGlobal *container.Split
	db          *sql.DB // объект базы данных приложения
}

// глобальные флаги
type globalFlags struct {
	beginFixAct bool
}

var g_v globalVariables // объект глобальных переменных
var g_f globalFlags

func App() {
	g_v.a = app.NewWithID("Main Journal")
	g_v.w = g_v.a.NewWindow("Main Journal v 0.0.1")
	g_v.w.Resize(fyne.NewSize(480, 300)) // размер окна
	g_v.w.SetFixedSize(true)
	g_v.w.CenterOnScreen()

	// СОЗДАЕМ ОБЩУЮ БАЗУ ДАННЫХ ПРИЛОЖЕНИЯ
	g_v.db = dataApp.CreateDB()
	defer g_v.db.Close()

	// БАГ - приложение ФИКСАКТ инициализируется МНОГО РАЗ!!! Нужно написать флаг
	rightContApp := fixact.NewApp(g_v.db) // В ДАЛЬНЕЙШЕМ ЗАМЕНИТ НА ЭКРАН ПРИВЕТСТВИЯ
	leftContApp := createSideBar()
	g_v.splitGlobal = container.NewHSplit(leftContApp, rightContApp)
	g_v.splitGlobal.SetOffset(0.01) // Устанавливаем разделитель в середину
	g_v.w.SetContent(g_v.splitGlobal)

	g_v.w.ShowAndRun()
}
