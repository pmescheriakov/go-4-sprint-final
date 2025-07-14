package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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
	strs := strings.Split(data, ",")
	if len(strs) != 3 {
		return 0, "", 0, errors.New("bad income data")
	}

	steps, activity, dur := strs[0], strs[1], strs[2]

	stepsCnt, err := strconv.Atoi(steps)
	if err != nil {
		return 0, "", 0, fmt.Errorf("bad income steps data: %w", err)
	}
	if stepsCnt <= 0 {
		return 0, "", 0, fmt.Errorf("bad income steps value: %v", stepsCnt)
	}

	duration, err := time.ParseDuration(dur)
	if err != nil {
		return 0, "", 0, fmt.Errorf("bad income duration data: %w", err)
	}
	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("bad income duration value: %v", duration)
	}

	return stepsCnt, activity, duration, nil
}

// distance
func distance(steps int, height float64) float64 {
	return (stepLengthCoefficient * height * float64(steps)) / mInKm
}

// meanSpeed
func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0.0
	}

	dist := distance(steps, height)

	return dist / duration.Hours()
}

// TrainingInfo
//
//	data — данные формата "3456,Ходьба,3h00m",
//	weight, height float64 — вес (кг.) и рост (м.) пользователя.
func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		return "", fmt.Errorf("bad parseTraining calculate: %w", err)
	}

	var calories float64
	var errSwitch error
	switch activity {
	case "Ходьба":
		calories, errSwitch = WalkingSpentCalories(steps, weight, height, duration)
		if errSwitch != nil {
			return "", fmt.Errorf("bad WalkingSpentCalories calculate: %w", err)
		}
	case "Бег":
		calories, errSwitch = RunningSpentCalories(steps, weight, height, duration)
		if errSwitch != nil {
			return "", fmt.Errorf("bad RunningSpentCalories calculate: %w", err)
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	dist := distance(steps, height)

	mSpeed := meanSpeed(steps, height, duration)

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			activity, duration.Hours(), dist, mSpeed, calories),
		nil
}

// RunningSpentCalories
//
//	steps — количество шагов,
//	weight, height — вес(кг.) и рост(м.) пользователя,
//	duration — продолжительность ходьбы.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0.0, errors.New("bad income data: some values <= 0")
	}

	mSpeed := meanSpeed(steps, height, duration)

	calories := duration.Minutes() * weight * mSpeed / minInH

	return calories, nil
}

// WalkingSpentCalories
//
//	steps — количество шагов,
//	weight, height — вес(кг.) и рост(м.) пользователя,
//	duration — продолжительность ходьбы.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= time.Duration(0) {
		return 0.0, errors.New("bad income data: some values <= 0")
	}

	mSpeed := meanSpeed(steps, height, duration)

	calories := walkingCaloriesCoefficient * (duration.Minutes() * weight * mSpeed / minInH)

	return calories, nil
}
