package settingapp

import (
	fyneview_utils "mainjournal/view/fyneview/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewApp(a fyne.App) *fyne.Container {

	nameApplicationLabel := widget.NewLabel("Setting applications")
	themeLabel := widget.NewLabel("🅃🄷🄴🄼🄴")

	buttLightTheme := widget.NewButton("Light", func() { fyneview_utils.SetLightTheme(a) })
	buttDarkTheme := widget.NewButton("Dark", func() { fyneview_utils.SetDarkTheme(a) })

	cont := container.NewVBox(nameApplicationLabel, themeLabel, buttLightTheme, buttDarkTheme)
	return cont
}
