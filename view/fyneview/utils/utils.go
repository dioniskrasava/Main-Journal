package fyneview_utils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

var containerSplitMainApp *container.Split

func SetLightTheme(a fyne.App) {
	a.Settings().SetTheme(theme.LightTheme())
}

func SetDarkTheme(a fyne.App) {
	a.Settings().SetTheme(theme.DarkTheme())
}

// получает глобальный объект окна приложения
// и переписывает местную переменную созданную
// для того, чтобы хранить объект главного окна
// и рисовать на нём всякие непристойности
// в других пакетах (которые к сожалению импортируются из главного view/fyneview и поэтому
// мы вынуждены работать через посредника)
func GetMainAppContainer(cont *container.Split) {
	containerSplitMainApp = cont
}

// side (r - для отрисовки на правой стороне, l - на левой)
// второй параметр - контейнер для отрисовки
func DrawMainAppCont(side string, cont *fyne.Container) {
	if side == "r" {
		containerSplitMainApp.Trailing = cont
		containerSplitMainApp.Refresh()
	} else if side == "l" {
		containerSplitMainApp.Leading = cont
		containerSplitMainApp.Refresh()
	}

}
