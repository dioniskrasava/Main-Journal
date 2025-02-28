package settingapp

import (
	fyneview_utils "mainjournal/view/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewApp(a fyne.App) *fyne.Container {

	nameApplicationLabel := widget.NewLabel("Настройки приложения \n\n\nТемы :")
	//themeLabel := widget.NewLabel("🅃🄷🄴🄼🄴")

	buttLightTheme := widget.NewButton("Светлая", func() { fyneview_utils.SetLightTheme(a) })
	buttDarkTheme := widget.NewButton("Тёмная", func() { fyneview_utils.SetDarkTheme(a) })
	contButThem := container.NewCenter(container.NewHBox(buttLightTheme, buttDarkTheme))

	labelLang := widget.NewLabel("Языки:")
	buttEng := widget.NewButton("Английский", func() { fyneview_utils.SetEnglishLangFixact() })
	buttRus := widget.NewButton("Русский", func() { fyneview_utils.SetRussianLangFixact() })
	contButLang := container.NewCenter(container.NewHBox(buttEng, buttRus))

	cont := container.NewVBox(nameApplicationLabel, contButThem, labelLang, contButLang)
	return cont
}
