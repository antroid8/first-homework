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

// Функция преобразовывает строку вводных данных в количество шагов, вид активности, продолжительность активности
func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	str := strings.Split(data, ",") // Преобразовали строку в слайс строк
	if len(str) != 3 {              // Проверили длину слайса
		return 0, "", 0, errors.New("не верное количество параметров")
	}
	steps, err := strconv.Atoi(str[0]) // Преобразовали первый элемент слайса в количество шагов
	if err != nil {                    // Проверили на наличие ошибок при преобразовании
		return 0, "", 0, errors.New("ошибка в парсинге")
	}
	if steps <= 0 { // Проверили количество шагов на положительность значения
		return 0, "", 0, errors.New("не верное количество шагов")
	}
	time, err := time.ParseDuration(str[2]) // Преобразовали второй элемент слайса в продолжительность активности
	if err != nil {                         // Проверили на наличие ошибок
		return 0, "", 0, errors.New("ошибка в парсинге")
	}
	if time <= 0 { // Проверили продолжительность активности на положительность значений
		return 0, "", 0, errors.New("не верная продолжительность занятий")
	}
	return steps, str[1], time, nil
}

// Функция находит пройденную дистанцию в километрах
func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	lengthStep := height * stepLengthCoefficient // Нашли длину шага
	distanceM := lengthStep * float64(steps)     // Нашли пройденную дистанцию в метрах
	distanceKm := distanceM / mInKm              // Перевели дистанцию в километры
	return distanceKm
}

// Функция находит среднюю скорость движения при активности
func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 { // Проверили продолжительность активности на положительность значения
		return 0
	}
	distanceKm := distance(steps, height)         // Получили пройденную дистанцию в километрах
	averageSpeed := distanceKm / duration.Hours() // Получили среднюю скорость
	return averageSpeed
}

// Функция выводит общую информацию об активности
func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activity, time, err := parseTraining(data) // Получили количество шагов, вид активности и ее продолжительность
	if err != nil {                                   //Проверили на наличие ошибок
		return "", fmt.Errorf("ошибка в парсинге")
	}
	switch activity { // С помощью switch перебрали возможные варианты активностей и в соответствии с ними вывели на экран информацию
	case "Бег", "бег":
		distanceRunning := distance(steps, height)                          //Нашли пройденную дистанцию
		averageSpeedRunning := meanSpeed(steps, height, time)               // Нашли среднюю скорость
		runningKkal, _ := RunningSpentCalories(steps, weight, height, time) // Нашли калории
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activity, time.Hours(), distanceRunning, averageSpeedRunning, runningKkal), nil
	case "Ходьба", "ходьба":
		distanceWalking := distance(steps, height)
		averageSpeedWalking := meanSpeed(steps, height, time)
		walkingKkal, _ := WalkingSpentCalories(steps, weight, height, time)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activity, time.Hours(), distanceWalking, averageSpeedWalking, walkingKkal), nil
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}

// Функция рассчитывает количество соженных калорий при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 { // Проверили вводные данные на положительность значений
		return 0, errors.New("отрицательное значение")
	}
	averageSpeed := meanSpeed(steps, height, duration)                 // Нашли среднюю скорость бега
	durationInMinutes := duration.Minutes()                            // Продолжительность бега в минутах
	aLotOfKkal := (weight * averageSpeed * durationInMinutes) / minInH // Количество затраченных калорий
	return aLotOfKkal, nil
}

// Функция рассчитывает количество соженных калорий при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 { // Проверили вводные данные на положительность значений
		return 0, errors.New("отрицательное значение")
	}
	averageSpeed := meanSpeed(steps, height, duration)                 // Нашли среднюю скорость ходьбы
	durationInMinutes := duration.Minutes()                            // Продолжительность ходьбы в минутах
	aLotOfKkal := (weight * averageSpeed * durationInMinutes) / minInH // Количество затраченных калорий при беге
	walkingKkal := aLotOfKkal * walkingCaloriesCoefficient             // Количество затраченных калорий при ходьбе
	return walkingKkal, nil
}
