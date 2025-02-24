package fyneview

import "fyne.io/fyne/v2/theme"

func setLightTheme() {
	g_v.a.Settings().SetTheme(theme.LightTheme())
}

func setDarkTheme() {
	g_v.a.Settings().SetTheme(theme.DarkTheme())
}
