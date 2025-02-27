package app2

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewApp() *fyne.Container {
	cont := container.NewHBox(widget.NewLabel("Lol"))
	return cont
}
