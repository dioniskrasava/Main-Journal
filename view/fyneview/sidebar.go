package fyneview

import (
	"fmt"
	"mainjournal/applications/fixact"
	stopwatch "mainjournal/applications/stpwtch"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// возвращает боковую панель главного приложения
// в виде контейнера файн
func createSideBar() *fyne.Container {
	but1 := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() { showFixact() })
	but2 := widget.NewButtonWithIcon("", theme.ContentClearIcon(), func() { showStopwatch() })

	cont := container.NewVBox(but1, but2)
	return cont
}

func showStopwatch() {
	if !g_f.beginStopwatch {
		fmt.Println("Первый запуск секундомера")
		g_f.beginStopwatch = true
		g_v.splitGlobal.Trailing = stopwatch.NewApp()
		g_v.w.Resize(fyne.NewSize(400, 110))
		g_v.splitGlobal.Refresh()
	} else {
		fmt.Println("Не первый запуск секундомера")
		g_v.splitGlobal.Trailing = stopwatch.NotNewApp()
		g_v.w.Resize(fyne.NewSize(400, 110))
		g_v.splitGlobal.Refresh()
	}

}

func showFixact() {
	g_v.splitGlobal.Trailing = fixact.NewApp()
	g_v.w.Resize(fyne.NewSize(480, 300))
	g_v.splitGlobal.Refresh()
}
