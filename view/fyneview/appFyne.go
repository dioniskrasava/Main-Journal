package fyneview

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func App() {
	a := app.NewWithID("Main Journal")
	w := a.NewWindow("Main Journal v 0.0.1")
	w.Resize(fyne.NewSize(800, 600)) // размер окна
	//w.SetFixedSize(true)
	w.CenterOnScreen()

	// Создаем два виджета, которые будут размещены в разделяемых панелях

	right := widget.NewLabel("Правая панель")

	but1 := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {})
	but2 := widget.NewButtonWithIcon("", theme.ContentClearIcon(), func() {})
	but3 := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {})
	but4 := widget.NewButtonWithIcon("", theme.ContentCutIcon(), func() {})
	but5 := widget.NewButtonWithIcon("", theme.ContentPasteIcon(), func() {})
	but6 := widget.NewButtonWithIcon("", theme.ContentRedoIcon(), func() {})
	but7 := widget.NewButtonWithIcon("", theme.ContentRemoveIcon(), func() {})
	but8 := widget.NewButtonWithIcon("", theme.ContentUndoIcon(), func() {})

	cont := container.NewVBox(but1, but2, but3, but4, but5, but6, but7, but8)

	// Создаем контейнер с разделяемыми панелями
	split := container.NewHSplit(cont, right)
	split.SetOffset(0.01) // Устанавливаем разделитель в середину

	// Устанавливаем контейнер в окно
	w.SetContent(split)

	w.ShowAndRun()
}
