package daysteps

import (
	"fmt"
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
	divorce := strings.Split(data, ",")
	if len(divorce) != 2 {
		return 0, 0, fmt.Errorf("неверные данные, длина слайса не равна 2")
	}
	firstElem := divorce[0]
	steps, err := strconv.Atoi(firstElem)
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка: %v", err)
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("количество шагов должно быть больше нуля!")
	}
	secondElem := divorce[1]
	duration, err := time.ParseDuration(secondElem)
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка: %v", err)
	}
	if duration <= 0 {
		return 0, 0, fmt.Errorf("продолжительность должна быть больше нуля!")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distanceM := float64(steps) * stepLength
	distanceKm := distanceM / mInKm
	calories := WalkingSpentCalories(steps, duration, weight, height)
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
		steps, distanceKm, calories)

	return result
}
