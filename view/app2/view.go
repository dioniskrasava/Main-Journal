package app2

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func CreateContentApp() *fyne.Container {

	activityType := widget.NewSelect([]string{"Книга", "Код", "Видео"}, func(value string) {})
	activityType.PlaceHolder = "Выбери активность"

	startTime := widget.NewEntry()
	endTime := widget.NewEntry()
	totalTime := widget.NewEntry()
	comment := widget.NewMultiLineEntry()
	addButton := widget.NewButton("Добавить активность", func() {})

	btnSupp0 := widget.NewButton("☰", func() {})
	btnSupp1 := widget.NewButton("🞴", func() {})
	btnSupp2 := widget.NewButton("🞴", func() {})
	btnSupp3 := widget.NewButton("🞴", func() {})

	lbTypeActivity := widget.NewLabel("Тип активности:")
	lbStartTime := widget.NewLabel("Время начала:")
	lbEndTime := widget.NewLabel("Время окончания:")
	lbTotalTime := widget.NewLabel("Общее время:")
	lbNameComment := widget.NewLabel("Комментарий:")

	row1 := container.NewGridWithColumns(3, lbTypeActivity, activityType, btnSupp0)
	row2 := container.NewGridWithColumns(3, lbStartTime, startTime, btnSupp1)
	row3 := container.NewGridWithColumns(3, lbEndTime, endTime, btnSupp2)
	row4 := container.NewGridWithColumns(3, lbTotalTime, totalTime, btnSupp3)
	row5 := container.NewGridWithColumns(2, lbNameComment, comment)
	row6 := container.NewGridWithColumns(1, addButton)

	globContainer := container.NewGridWithRows(6, row1, row2, row3, row4, row5, row6)

	return globContainer
}

func CreateContentApp2() *fyne.Container {
	// Создаем виджеты
	label1 := widget.NewLabel("Label 1")
	label2 := widget.NewLabel("Label 2")
	label3 := widget.NewLabel("Label 3")
	button := widget.NewButton("Button", func() {})

	// Создаем контейнер с GridLayout
	grid := container.New(layout.NewGridLayout(3),
		label1, layout.NewSpacer(), label2, // Первая строка
		layout.NewSpacer(), button, layout.NewSpacer(), // Вторая строка
		label3, layout.NewSpacer(), layout.NewSpacer(), // Третья строка
	)

	return grid
}
