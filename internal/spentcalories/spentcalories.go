package spentcalories

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию

	return 0, "", time.Duration(1), nil
}

// distance
func distance(steps int, height float64) float64 {
	return (stepLengthCoefficient * height * float64(steps)) / mInKm
}

// meanSpeed
func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration < 0 {
		return 0.
	}

	dist := distance(steps, height)

	return dist / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию

	return "", nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию

	return 0.0, nil
}

// WalkingSpentCalories
//
//	steps — количество шагов,
//	weight, height — вес(кг.) и рост(м.) пользователя,
//	duration — продолжительность ходьбы.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps < 0 || weight < 0 || height < 0 || duration < time.Duration(0) {
		return 0.0, errors.New("bad income data")
	}

	mSpeed := meanSpeed(steps, height, duration)

	calories := walkingCaloriesCoefficient * (duration.Minutes() * weight * mSpeed / 60)

	return calories, nil
}
