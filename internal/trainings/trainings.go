package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// создайте структуру Training
// Структура Training описывает содержит все необходимые данные о тренировке: количество шагов,
// тип тренировки, длительность тренировки, а также данные из структуры personaldata.Personal
type Training struct {
	Steps                 int           // Количество шагов, проделанных за тренировку.
	TrainingType          string        // Тип тренировки(бег или ходьба).
	Duration              time.Duration // Длительность тренировки.
	personaldata.Personal               // Встраиваем существующую структуру Personal из пакета "personaldata"
}

// создайте метод Parse()
// Метод Parse() принимает строку с данными которая содержит кол-во шагов, вид активности
// и продолжительность активности. Метод записывает данные в структуру Training
// возвращает ошибку. При этом все преобразуется в нужные типы.
func (t *Training) Parse(datastring string) (err error) {

	// Делим входящую строку на части, делаем слайсы.
	separatePart := strings.Split(datastring, ",")
	if len(separatePart) != 3 {
		return fmt.Errorf("invalid data format")
	}

	// Преобразуем строку в число, сохраняем полученное значение в структуру
	countSteps, err := strconv.Atoi(separatePart[0])
	if err != nil {
		return fmt.Errorf("invalid data format for steps: %w", err)
	}
	t.Steps = countSteps
	// Проверяем второй элемент слайса на соответствие известных нам
	// типов тренировок (Бег или Ходьба)
	if separatePart[1] != "Бег" && separatePart[1] != "Ходьба" {
		return fmt.Errorf("invalid data format for activityType")
	}
	t.TrainingType = separatePart[1] // Присваиваем значение в поле структуры

	// Преобразуем строку в time.Duration, сохраняем полученное значение в структуру.
	duration, err := time.ParseDuration(separatePart[2])
	if err != nil {
		return fmt.Errorf("invalid data format for duration: %w", err)
	}
	t.Duration = duration

	return nil
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() string {

	// Рассчитываем дистанцию и среднюю скорость, используя функцию из пакета spentenergy.
	dist := spentenergy.Distance(t.Steps)
	speed := spentenergy.MeanSpeed(t.Steps, t.Duration)

	var calories float64

	// Определяем тип тренировки и считаем калории для каждого из них.
	switch t.TrainingType {
	case "Бег":
		calories = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Duration)
	case "Ходьба":
		calories = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	default:
		return "invalid data format for activityType"
	}
	// Формируем строку результата
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		t.TrainingType, t.Duration.Hours(), dist, speed, calories)
}
