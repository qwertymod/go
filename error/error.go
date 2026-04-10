package error

import (
	"github.com/fatih/color"
)

func PrintError(value any) { // interface{}
	//color.Red(value) - так нельзя сделать из за разности типов

	switch t := value.(type) {
	case string:
		color.Red(t)
	case int:
		color.Red("Код ошибки %d", t)
	case error:
		color.Red(t.Error())
	default:
		color.Red("Неизвестный тип ошибки")
	}
}