package fyneview

import (
	"mainjournal/applications/fixact"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

// глобальные переменные
type globalVariables struct {
	w           fyne.Window
	splitGlobal *container.Split
}

// глобальные флаги
type globalFlags struct {
	beginStopwatch bool
}

var g_v globalVariables
var g_f globalFlags

func App() {
	a := app.NewWithID("Main Journal")
	g_v.w = a.NewWindow("Main Journal v 0.0.1")
	g_v.w.Resize(fyne.NewSize(480, 300)) // размер окна
	//w.SetFixedSize(true)
	g_v.w.CenterOnScreen()

	// Создаем два виджета, которые будут размещены в разделяемых панелях

	right := fixact.NewApp()

	cont := createSideBar()

	// Создаем контейнер с разделяемыми панелями
	g_v.splitGlobal = container.NewHSplit(cont, right)
	g_v.splitGlobal.SetOffset(0.01) // Устанавливаем разделитель в середину

	// Устанавливаем контейнер в окно
	g_v.w.SetContent(g_v.splitGlobal)

	g_v.w.ShowAndRun()
}
