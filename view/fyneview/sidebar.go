package fyneview

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// возвращает боковую панель главного приложения
// в виде контейнера файн
func createSideBar() *fyne.Container {
	but1 := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {})
	but2 := widget.NewButtonWithIcon("", theme.ContentClearIcon(), func() {})
	but3 := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {})
	but4 := widget.NewButtonWithIcon("", theme.ContentCutIcon(), func() {})
	but5 := widget.NewButtonWithIcon("", theme.ContentPasteIcon(), func() {})
	but6 := widget.NewButtonWithIcon("", theme.ContentRedoIcon(), func() {})
	but7 := widget.NewButtonWithIcon("", theme.ContentRemoveIcon(), func() {})
	but8 := widget.NewButtonWithIcon("", theme.ContentUndoIcon(), func() {})

	cont := container.NewVBox(but1, but2, but3, but4, but5, but6, but7, but8)
	return cont
}
