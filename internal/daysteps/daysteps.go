package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

// Функция для преобразования вводных данных
func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	v := strings.Split(data, ",") // Из строки делаем слайс строк
	if len(v) != 2 {              // Проверяем слайс на нужное нам количество элементов
		return 0, 0, fmt.Errorf("не верное количество введенных данных")
	}
	steps, err := strconv.Atoi(v[0]) // Преобразовываем первый элементы слайса строк в количество шагов
	if err != nil {                  // Проверяем на наличие ошибок при преобразовании строки в число
		return 0, 0, fmt.Errorf("не верный тип количества шагов")
	}
	if steps <= 0 { // Проверяем количество шагов на положительность значения
		return 0, 0, fmt.Errorf("количество шагов должно быть больше 0")
	}
	time, err := time.ParseDuration(v[1]) // Преобразовываем второй элемент слайса строк во время
	if err != nil {                       // Проверяем на наличие ошибок при преобразовании
		log.Println(err)
		return 0, 0, fmt.Errorf("не верный тип времени")
	}
	if time <= 0 { // Проверяем время на положительность значения
		log.Println(err)
		return 0, 0, fmt.Errorf("отрицательное значение времени")
	}
	return steps, time, nil
}

// Функция для подсчета активности и сожженных калорий
func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, time, err := parsePackage(data) // Вычленили количество шагов, продолжительность активности из вводных данных
	if err != nil {                        // Проверили на наличие ошибок при преобразовании
		log.Println("некорректный формат", err)
		return ""
	}
	if steps <= 0 { // Проверили количество шагов на положительность значения
		log.Println("количество шагов должно быть больше 0", err)
		return ""
	}
	if time <= 0 { // Проверили продолжительность активности на положительность значения
		log.Println("не может быть отрицательного времени", err)
		return ""
	}
	distanceM := float64(steps) * stepLength                                   // Получили пройденную дистанцию в метрах
	distanceKm := distanceM / mInKm                                            // Дистанция в километрах
	kkal, _ := spentcalories.WalkingSpentCalories(steps, weight, height, time) // Получили количество сожженных калорий
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanceKm, kkal)
	return result
}
