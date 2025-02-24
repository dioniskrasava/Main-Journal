package fyneview

import (
	"mainjournal/applications/fixact"
	"mainjournal/applications/stopwatch"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// возвращает боковую панель главного приложения
// в виде контейнера файн
func createSideBar() *fyne.Container {
	fixactButt := widget.NewButton("🅵", func() { showFixact() })
	stopwatButt := widget.NewButton("🆂", func() { showStopwatch() })

	labelSepar := widget.NewLabel("🅃\n🄷\n🄴\n🄼\n🄴")

	butLThem := widget.NewButton("🅻", func() { setLightTheme() }) // Light theme
	butDThem := widget.NewButton("🅳", func() { setDarkTheme() })  // Dark theme

	cont := container.NewVBox(fixactButt, stopwatButt, labelSepar, butLThem, butDThem)
	return cont
}

func showStopwatch() {
	g_v.splitGlobal.Trailing = stopwatch.NewApp()
	g_v.w.Resize(fyne.NewSize(480, 300))
	g_v.w.SetFixedSize(false)
	g_v.splitGlobal.Refresh()
}

func showFixact() {
	g_v.splitGlobal.Trailing = fixact.NewApp(g_v.db)
	g_v.w.Resize(fyne.NewSize(480, 300))
	g_v.w.SetFixedSize(true)
	g_v.splitGlobal.Refresh()
}

func setLightTheme() {
	g_v.a.Settings().SetTheme(theme.LightTheme())
}

func setDarkTheme() {
	g_v.a.Settings().SetTheme(theme.DarkTheme())
}
