package fyneview_utils

import "fyne.io/fyne/v2/widget"

// СТРУКТУРА ВИДЖЕТОВ     F I X A C T
// для того, чтобы иметь возможность получать их
// и изменять язык приложения

type WidgetsFixact struct {
	ActivityType      *widget.Select
	StartTime         *widget.Entry
	EndTime           *widget.Entry
	TotalTime         *widget.Entry
	Comment           *widget.Entry
	AddButton         *widget.Button
	LabelTypeActivity *widget.Label
	LabelStartTime    *widget.Label
	LabelEndTime      *widget.Label
	LabelTotalTime    *widget.Label
	LabelNameComment  *widget.Label
}

// создаем пустой объект виджетов микроприложения Фиксакт
var widgetsFixact WidgetsFixact
