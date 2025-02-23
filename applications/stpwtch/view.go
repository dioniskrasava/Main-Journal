package stopwatch

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// инициализация НАЧАЛЬНОГО СОСТОЯНИЯ секундомера
func widgets_begin() *fyne.Container {

	bgColor := theme.BackgroundColor()
	rgbaColor := color.RGBAModel.Convert(bgColor).(color.RGBA)

	var input *canvas.Text

	// ЕСЛИ ЦВЕТ ФОНА БЕЛЫЙ
	if rgbaColor == (color.RGBA{255, 255, 255, 255}) {
		input = canvas.NewText("00:00:00:000", color.RGBA{0, 0, 0, 222})
	} else {
		input = canvas.NewText("00:00:00:000", color.RGBA{255, 255, 255, 255})
	}

	input.TextSize = 35 // Увеличение размера шрифта

	btn_str := widget.NewButton("       Start       ", nil)
	btn_pause := widget.NewButton("      Pause      ", nil)
	btn_clear := widget.NewButton("      Clear      ", nil)

	wdgts.btn_str = btn_str
	wdgts.btn_pause = btn_pause
	wdgts.btn_clear = btn_clear
	wdgts.input = input

	// изначально кнопки паузы и чистки - неактивны
	btn_pause.Disable()
	btn_clear.Disable()

	btn_str.OnTapped = logic_timer("start", &flags, wdgts, &time_val)
	btn_pause.OnTapped = logic_timer("pause", &flags, wdgts, &time_val)
	btn_clear.OnTapped = logic_timer("clear", &flags, wdgts, &time_val)

	wdgts.btn_cont = container.NewHBox(btn_str, btn_pause, btn_clear)

	wdgts.cont_begin = container.NewVBox(container.NewCenter(input), container.NewCenter(wdgts.btn_cont))

	containerApp := wdgts.cont_begin
	return containerApp
}

func widgetsNot_begin() *fyne.Container {

	timersTime := time_val.times

	bgColor := theme.BackgroundColor()
	rgbaColor := color.RGBAModel.Convert(bgColor).(color.RGBA)

	var input *canvas.Text

	// ЕСЛИ ЦВЕТ ФОНА БЕЛЫЙ
	if rgbaColor == (color.RGBA{255, 255, 255, 255}) {
		input = canvas.NewText(timersTime, color.RGBA{0, 0, 0, 222})
	} else {
		input = canvas.NewText(timersTime, color.RGBA{255, 255, 255, 255})
	}

	input.TextSize = 35 // Увеличение размера шрифта

	btn_str := widget.NewButton("       Start       ", nil)
	btn_pause := widget.NewButton("      Pause      ", nil)
	btn_clear := widget.NewButton("      Clear      ", nil)

	wdgts.btn_str = btn_str
	wdgts.btn_pause = btn_pause
	wdgts.btn_clear = btn_clear
	wdgts.input = input

	// изначально кнопки паузы и чистки - неактивны
	btn_pause.Disable()
	btn_clear.Disable()

	btn_str.OnTapped = logic_timer("start", &flags, wdgts, &time_val)
	btn_pause.OnTapped = logic_timer("pause", &flags, wdgts, &time_val)
	btn_clear.OnTapped = logic_timer("clear", &flags, wdgts, &time_val)

	wdgts.btn_cont = container.NewHBox(btn_str, btn_pause, btn_clear)

	wdgts.cont_begin = container.NewVBox(container.NewCenter(input), container.NewCenter(wdgts.btn_cont))

	containerApp := wdgts.cont_begin
	return containerApp
}
