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
	// TODO: реализовать функцию
	str := strings.Split(data, ",")
	if len(str) != 3 {
		return 0, "", 0, nil
	}
	steps, err := strconv.Atoi(str[0])
	if err != nil {
		return 0, "", 0, err
	}
	time, err := time.ParseDuration(str[2])
	if err != nil {
		return 0, "", 0, err
	}
	return steps, str[1], time, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	lengthStep := height * stepLengthCoefficient
	distanceM := lengthStep * float64(steps)
	distanceKm := distanceM / mInKm
	return distanceKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	distanceKm := distance(steps, height)
	averageSpeed := distanceKm / duration.Hours()
	return averageSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activity, time, err := parseTraining(data)
	if err != nil {
		return "", err
	}
	switch activity {
	case "Бег", "бег":
		distanceRunning := distance(steps, height)
		averageSpeedRunning := meanSpeed(steps, height, time)
		runningKkal, _ := RunningSpentCalories(steps, weight, height, time)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %s ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий %.2f\n", activity, time, distanceRunning, averageSpeedRunning, runningKkal), nil
	case "Ходьба", "ходьба":
		distanceWalking, _ := WalkingSpentCalories(steps, weight, height, time)
		averageSpeedWalking := meanSpeed(steps, height, time)
		walkingKkal, _ := RunningSpentCalories(steps, weight, height, time)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %s ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий %.2f\n", activity, time, distanceWalking, averageSpeedWalking, walkingKkal), nil
	default:
		return fmt.Sprintln("неизвестный тип тренировки"), nil
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("incorrect data")
	}
	averageSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	aLotOfKkal := (weight * averageSpeed * durationInMinutes) / minInH
	return aLotOfKkal, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("incorrect data")
	}
	averageSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	aLotOfKkal := (weight * averageSpeed * durationInMinutes) / minInH
	walkingKkal := aLotOfKkal * walkingCaloriesCoefficient
	return walkingKkal, nil
}
