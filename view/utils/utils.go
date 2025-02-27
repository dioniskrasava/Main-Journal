package fyneview_utils

import (
	"fmt"

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
func ExportMainAppContainer(cont *container.Split) {
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

func SetEnglishLangFixact() {
	fmt.Println("ENG")

	widgetsFixact.ActivityType.PlaceHolder = "Select activites"

	widgetsFixact.LabelTypeActivity.SetText("Type activites : ")
	widgetsFixact.LabelStartTime.SetText("Start time : ")
	widgetsFixact.LabelEndTime.SetText("End time : ")
	widgetsFixact.LabelTotalTime.SetText("Total time : ")
	widgetsFixact.LabelNameComment.SetText("Comment : ")

	// НУЖНО СКИНУТЬ В ПРОКЛАДКУ СТРУКТУРУ ВИДЖЕТОВ ФИКСАКТА
}

func ExportWidgetsFixact(widgets WidgetsFixact) {
	widgetsFixact = widgets // записываем полученные виджеты приложения в глоб.перем.
}
