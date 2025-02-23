package stopwatch

import (
	"fyne.io/fyne/v2"
)

// Приложение секундомера
func NewApp() *fyne.Container {
	contenApp := widgets_begin()
	return contenApp
}

func NotNewApp() *fyne.Container {
	contenApp := widgetsNot_begin()
	return contenApp

}
