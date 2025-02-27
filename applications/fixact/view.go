package fixact

import (
	"database/sql"
	"fmt"
	"mainjournal/applications/fixact/pos"
	fyneview_utils "mainjournal/view/fyneview/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

func createInterfaceApp(db *sql.DB) (content *fyne.Container) {
	// Элементы интерфейса
	activityType := widget.NewSelect([]string{"Книга", "Код", "Видео"}, func(value string) { activityTypeEvent(value) })
	activityType.PlaceHolder = "Выбери активность"

	startTime := widget.NewEntry()
	endTime := widget.NewEntry()
	totalTime := widget.NewEntry()
	comment := widget.NewMultiLineEntry()
	addButton := widget.NewButton("Добавить активность", func() { addAct(widgtsApp, db) })

	btnSupp0 := widget.NewButton("☰", func() { drawTheSettingsWindow() })
	btnSupp1 := widget.NewButton("🞴", func() { btnSupp1Event(widgtsApp) })
	btnSupp2 := widget.NewButton("🞴", func() { btnSupp2Event(widgtsApp) })
	btnSupp3 := widget.NewButton("🞴", func() { btnSupp3Event(widgtsApp) })
	comment.OnChanged = func(text string) { commentFielChanged(text) }

	globContainer := container.NewWithoutLayout()

	h1 := float32(10)
	h2 := float32(50)
	h3 := float32(90)
	h4 := float32(130)
	h5 := float32(170)
	h6 := float32(250)

	labelTypeActivity := widget.NewLabel("Тип активности:")
	labelStartTime := widget.NewLabel("Время начала:")
	labelEndTime := widget.NewLabel("Время окончания:")
	labelTotalTime := widget.NewLabel("Общее время:")
	labelNameComment := widget.NewLabel("Комментарий:")

	// ДУБЛИРУЕМ СОЗДАННЫЕ ВИДЖЕТЫ В ГЛОБАЛЬНУЮ ПЕРЕМЕННУЮ
	widgtsApp = fyneview_utils.WidgetsFixact{
		ActivityType:      activityType,
		StartTime:         startTime,
		EndTime:           endTime,
		TotalTime:         totalTime,
		Comment:           comment,
		AddButton:         addButton,
		LabelTypeActivity: labelTypeActivity,
		LabelStartTime:    labelStartTime,
		LabelEndTime:      labelEndTime,
		LabelTotalTime:    labelTotalTime,
		LabelNameComment:  labelNameComment,
	}

	fyneview_utils.ExportWidgetsFixact(widgtsApp)

	pos.AddRow(globContainer, WIDTH, h1, labelTypeActivity, activityType, btnSupp0)
	pos.AddRow(globContainer, WIDTH, h2, labelStartTime, startTime, btnSupp1)
	pos.AddRow(globContainer, WIDTH, h3, labelEndTime, endTime, btnSupp2)
	pos.AddRow(globContainer, WIDTH, h4, labelTotalTime, totalTime, btnSupp3)
	pos.AddRow(globContainer, WIDTH, h5, labelNameComment, comment)
	pos.AddRow(globContainer, WIDTH, h6, widget.NewLabel(""), addButton)

	return globContainer
}

// возвращаем уже созданный интерфейс ФИКСАКТА
// НУЖЕН В ЦЕЛЯХ оптимизации (зачем каждый раз при отрисовке - создавать новый объект?)
func createOldInterfaceApp(db *sql.DB) *fyne.Container {

	btnSupp0 := widget.NewButton("☰", func() { drawTheSettingsWindow() })
	btnSupp1 := widget.NewButton("🞴", func() { btnSupp1Event(widgtsApp) })
	btnSupp2 := widget.NewButton("🞴", func() { btnSupp2Event(widgtsApp) })
	btnSupp3 := widget.NewButton("🞴", func() { btnSupp3Event(widgtsApp) })
	addButton := widget.NewButton("Добавить активность", func() { addAct(widgtsApp, db) })

	globContainer := container.NewWithoutLayout()

	h1 := float32(10)
	h2 := float32(50)
	h3 := float32(90)
	h4 := float32(130)
	h5 := float32(170)
	h6 := float32(250)

	pos.AddRow(globContainer, WIDTH, h1, widgtsApp.LabelTypeActivity, widgtsApp.ActivityType, btnSupp0)
	pos.AddRow(globContainer, WIDTH, h2, widgtsApp.LabelStartTime, widgtsApp.StartTime, btnSupp1)
	pos.AddRow(globContainer, WIDTH, h3, widgtsApp.LabelEndTime, widgtsApp.EndTime, btnSupp2)
	pos.AddRow(globContainer, WIDTH, h4, widgtsApp.LabelTotalTime, widgtsApp.TotalTime, btnSupp3)
	pos.AddRow(globContainer, WIDTH, h5, widgtsApp.LabelNameComment, widgtsApp.Comment)
	pos.AddRow(globContainer, WIDTH, h6, widget.NewLabel(""), addButton)

	return globContainer
}

// activityType EVENTS (CLCIK)
func activityTypeEvent(value string) {
	fmt.Println("Выбран ", value)
}

// ФУНКЦИИ ДЛЯ РАБОТЫ ВСПОМОГАТЕЛЬНЫХ КНОПОК
// ДЛЯ НИХ НУЖНО ПИСАТЬ ГЛОБАЛЬНЫЙ ОБЪЕКТ ВИДЖЕТОВ ПРИЛОЕЖНИЯ, ЧТОБЫ В ЭТИХ ФУНКЦИЯХ
// СЧИТЫВАТЬ ЗНАЧЕНИЯ С ЭТИХ ВИДЖЕТОВ И ПОМЕЩАТЬ ИХ В AFixD (объект состояний полей ввода ФИКС-АКТА)
func btnSupp1Event(widgtsApp fyneview_utils.WidgetsFixact) {
	widgtsApp.StartTime.SetText(getNow())
	appFieldVal = readFieldApp(widgtsApp)
	//fmt.Println(appFieldVal)

}

func btnSupp2Event(widgtsApp fyneview_utils.WidgetsFixact) {
	widgtsApp.EndTime.SetText(getNow())
	appFieldVal = readFieldApp(widgtsApp)
	//fmt.Println(appFieldVal)
}

func btnSupp3Event(widgtsApp fyneview_utils.WidgetsFixact) {
	widgtsApp.TotalTime.SetText(getActTime(widgtsApp.EndTime.Text, widgtsApp.StartTime.Text))
	appFieldVal = readFieldApp(widgtsApp)
	//fmt.Println(appFieldVal)
}

// вызывается если пользователь что-то написал в поле комментариев
func commentFielChanged(text string) {
	fmt.Printf("Text changed: %s\n", text)
	appFieldVal = readFieldApp(widgtsApp)
	//fmt.Println(appFieldVal)
}
