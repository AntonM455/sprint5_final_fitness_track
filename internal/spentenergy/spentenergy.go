package spentenergy

import "time"

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
	speed     = 1.39  // средняя скорость в м/с
)

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
// duration time.Duration — длительность тренировки.
//
// Создайте функцию ниже.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	// вызываем функцию meanSpeed, чтобы посчитать среднюю скорость.
	walkMeanSpead := MeanSpeed(steps, duration)
	// Возвращаем кол-во потраченных калорий при ходьбе.
	return ((walkingCaloriesWeightMultiplier * weight) + (walkMeanSpead*walkMeanSpead/height)*walkingSpeedHeightMultiplier) * duration.Hours() * minInH

}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.
//
// Создайте функцию ниже.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	// вызываем функцию meanSpeed, чтобы посчитать среднюю скорость
	runMeanSpead := MeanSpeed(steps, duration)
	// Возвращаем кол-во потраченных калорий при беге
	return ((runningCaloriesMeanSpeedMultiplier * runMeanSpead) - runningCaloriesMeanSpeedShift) * weight
}

// МeanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
//
// Создайте функцию ниже.
func MeanSpeed(steps int, duration time.Duration) float64 {
	if duration <= 0 { // Проверяем, чтобы продолжительность не была равна 0
		return 0
	}
	return Distance(steps) / duration.Hours() // Возвращаем среднюю скорость
}

// Distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Для расчета дистанции нужно шаги умножить на длину шага lenStep и разделить на mInKm
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
//
// Создайте функцию ниже
func Distance(steps int) float64 {
	distanceInKm := float64(steps) * lenStep / float64(mInKm)
	return distanceInKm
}
