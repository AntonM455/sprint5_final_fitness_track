package personaldata

import "fmt"

// Ниже создайте структуру Personal
// Структура описывает пользователя и используется в других пакетах
type Personal struct {
	Name   string  // Имя рользователя
	Weight float64 // Вес пользователя
	Height float64 // Рост пользователя
}

// Ниже создайте метод Print()
// Метод Print() позволяет вывести на экран всю информацию о пользователе
func (p Personal) Print() {
	fmt.Printf(
		"Имя пользователя: %s\nВес пользователя: %.2f\nРост пользователя: %.2f\n",
		p.Name, p.Weight, p.Height,
	)
}
