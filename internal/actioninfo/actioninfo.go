package actioninfo

import (
	"fmt"
)

// создайте интерфейс DataParser
type DataParser interface {
	Parse(datastring string) error // Распарсить строку с данными о шагах, типе тренировки
	// и продолжительности
	ActionInfo() string // Данные о прогулке
}

// создайте функцию Info()
// функция Info() принимает слайс строк с данными о тренировках или прогулках и выводит
// информацию об активности с помощью метода ActionInfo().
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			fmt.Println("invalid data format:", err)
			continue
		}
		fmt.Println(dp.ActionInfo())
	}
}
