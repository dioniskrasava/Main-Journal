package fixact

import "log"

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
