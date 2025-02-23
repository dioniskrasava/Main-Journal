package fyneview

import (
	"mainjournal/applications/fixact"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func App() {
	a := app.NewWithID("Main Journal")
	w := a.NewWindow("Main Journal v 0.0.1")
	w.Resize(fyne.NewSize(800, 600)) // размер окна
	//w.SetFixedSize(true)
	w.CenterOnScreen()

	// Создаем два виджета, которые будут размещены в разделяемых панелях

	right := fixact.NewApp()

	cont := createSideBar()

	// Создаем контейнер с разделяемыми панелями
	split := container.NewHSplit(cont, right)
	split.SetOffset(0.01) // Устанавливаем разделитель в середину

	// Устанавливаем контейнер в окно
	w.SetContent(split)

	w.ShowAndRun()
}
