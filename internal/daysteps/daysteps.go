package daysteps

import (
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	v := strings.Split(data, ",")
	if len(v) != 2 {
		return 0, 0, nil
	}
	steps, err := strconv.Atoi(v[0])
	if err != nil || steps <= 0 {
		return 0, 0, err
	}
	time, err := time.ParseDuration(v[1])
	if err != nil {
		return 0, 0, err
	}
	return steps, time, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
}
