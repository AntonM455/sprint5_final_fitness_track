package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

const (
	StepLength = 0.65
)

// создайте структуру DaySteps
// структура DaySteps содержит все необходимые данные о дневных прогулках: количество шагов,
// длительность, а также данные из структуры personaldata.Personal, то есть имя, вес и
// рост пользователя.
type DaySteps struct {
	Steps                 int           // Количество шагов, проделанных за тренировку.
	Duration              time.Duration // Длительность тренировки.
	personaldata.Personal               // Встраиваем существующую структуру Personal из пакета "personaldata"
}

// создайте метод Parse()
// метод Parse() парсит строку с данными формата "678,0h50m"
// и записывает данные в соответствующие поля структуры DaySteps.
func (ds *DaySteps) Parse(datastring string) (err error) {
	// Делим входящую строку на части, делаем слайсы.
	separatePart := strings.Split(datastring, ",")
	if len(separatePart) != 2 {
		return fmt.Errorf("invalid data format")
	}
	// Преобразуем строку в число, сохраняем полученное значение в структуру кол-ва шагов.
	countSteps, err := strconv.Atoi(separatePart[0])
	if err != nil {
		return fmt.Errorf("invalid data format for steps: %w", err)
	}
	ds.Steps = countSteps // Присваиваем значение в поле структуры

	// Преобразуем строку в time.Duration, сохраняем полученное значение в структуру.
	duration, err := time.ParseDuration(separatePart[1])
	if err != nil {
		return fmt.Errorf("invalid data format for duration: %w", err)
	}
	ds.Duration = duration
	return nil
}

// создайте метод ActionInfo()
// метод ActionInfo() формирует и возвращает строку с данными о прогулке
func (ds DaySteps) ActionInfo() string {

	// Рассчитываем дистанцию и среднюю скорость, используя функцию из пакета spentenergy.
	distanceInMt := float64(ds.Steps) * StepLength
	distanceInKm := distanceInMt / 1000
	caloriesBurned := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	// Формируем строку результата
	return fmt.Sprintf("Количество шагов: %d. \nДистанция составила %.2f км. \nВы сожгли %.2f ккал.",
		ds.Steps, distanceInKm, caloriesBurned)
}
