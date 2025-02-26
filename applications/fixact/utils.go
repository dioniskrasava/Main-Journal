package fixact

import (
	"fmt"
	"log"
	fyneview_utils "mainjournal/view/fyneview/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// возвращает значения полей приложения
// в виде кастомной структуры Activity
func readFieldApp(w Widgets) Activity {
	activity := Activity{
		Type:      w.activityType.Selected,
		StartTime: w.startTime.Text,
		EndTime:   w.endTime.Text,
		TotalTime: w.totalTime.Text,
		Comment:   w.comment.Text,
	}
	return activity
}

// очистить поля приложения
// вызывается после добавления активности в базу
func cleanFielApp(w Widgets) {
	w.activityType.SetSelected("")
	w.startTime.SetText("")
	w.endTime.SetText("")
	w.totalTime.SetText("")
	w.comment.SetText("")
}

// записывает значения глобальной переменной
// в поля приложения
func writeFieldApp(w Widgets) {
	log.Println(appFieldVal)
	w.activityType.SetSelected(appFieldVal.Type)
	w.startTime.SetText(appFieldVal.StartTime)
	w.endTime.SetText(appFieldVal.EndTime)
	w.totalTime.SetText(appFieldVal.TotalTime)
	w.comment.SetText(appFieldVal.Comment)
}

func drawTheSettingsWindow() {
	buttBack := widget.NewButton("BACK", func() { drawFixActOnMainWindow() })
	cont := container.NewVBox(widget.NewLabel("Заработало!"), buttBack)
	fyneview_utils.DrawMainAppCont("r", cont)
}

func drawFixActOnMainWindow() {
	contFixAct := OldApp(maindb)

	// Проверяем, что contFixAct является *fyne.Container
	if cont, ok := contFixAct.(*fyne.Container); ok {
		fyneview_utils.DrawMainAppCont("r", cont)
	} else {
		// Обработка ошибки, если contFixAct не является *fyne.Container
		fmt.Println("Ошибка: contFixAct не является *fyne.Container")
	}
}
