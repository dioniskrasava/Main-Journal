package fixact

import (
	"fmt"
	fyneview_utils "mainjournal/view/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// возвращает значения полей приложения
// в виде кастомной структуры Activity
func readFieldApp(w fyneview_utils.WidgetsFixact) Activity {
	activity := Activity{
		Type:      w.ActivityType.Selected,
		StartTime: w.StartTime.Text,
		EndTime:   w.EndTime.Text,
		TotalTime: w.TotalTime.Text,
		Comment:   w.Comment.Text,
	}
	return activity
}

// очистить поля приложения
// вызывается после добавления активности в базу
func cleanFielApp(w fyneview_utils.WidgetsFixact) {
	w.ActivityType.SetSelected("")
	w.StartTime.SetText("")
	w.EndTime.SetText("")
	w.TotalTime.SetText("")
	w.Comment.SetText("")
}

// записывает значения глобальной переменной
// в поля приложения
func writeFieldApp(w fyneview_utils.WidgetsFixact) {
	//log.Println(appFieldVal) // для теста
	w.ActivityType.SetSelected(appFieldVal.Type)
	w.StartTime.SetText(appFieldVal.StartTime)
	w.EndTime.SetText(appFieldVal.EndTime)
	w.TotalTime.SetText(appFieldVal.TotalTime)
	w.Comment.SetText(appFieldVal.Comment)
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
