package fyneview

import (
	"mainjournal/view/app2"
	"mainjournal/view/fixact"
	settingapp "mainjournal/view/setting_app"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// возвращает боковую панель главного приложения
// в виде контейнера файн
func createSideBar() *fyne.Container {
	fixactButt := widget.NewButton("🅵", func() { showFixact() })
	stopwatButt := widget.NewButton("🆂", func() { showApp2() })

	// прокладка - для того чтобы кнопку настроек сделать внизу
	labelSepar := widget.NewLabel("\n\n\n\n\n\n")

	settingsAppButt := widget.NewButton("☰", func() { showSettingsApp() })
	//▤ ℵ ▢ ☰
	cont := container.NewVBox(fixactButt, stopwatButt, labelSepar, settingsAppButt)
	return cont
}

func showApp2() {
	g_v.splitGlobal.Trailing = app2.CreateContentApp2()
	g_v.w.Resize(fyne.NewSize(480, 300))
	//g_v.w.SetFixedSize(false)
	g_v.splitGlobal.Refresh()
}

func showFixact() {
	if !g_f.beginFixAct {
		//log.Println("ОТРИСОВАЛИ НОВОЕ ОКНО")
		g_v.splitGlobal.Trailing = fixact.NewApp(g_v.db)
		g_f.beginFixAct = true
	} else {
		//log.Println("ОТРИСОВАЛИ СТАРОЕ ОКНО")
		g_v.splitGlobal.Trailing = fixact.OldApp(g_v.db)
	}

	g_v.w.Resize(fyne.NewSize(480, 300))
	g_v.w.SetFixedSize(true)
	g_v.splitGlobal.Refresh()
}

func showSettingsApp() {
	g_v.splitGlobal.Trailing = settingapp.NewApp(g_v.a)
	g_v.w.Resize(fyne.NewSize(480, 300))
	//g_v.w.SetFixedSize(false)
	g_v.splitGlobal.Refresh()
}
