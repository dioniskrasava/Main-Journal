package fyneview_utils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

func SetLightTheme(a fyne.App) {
	a.Settings().SetTheme(theme.LightTheme())
}

func SetDarkTheme(a fyne.App) {
	a.Settings().SetTheme(theme.DarkTheme())
}
