package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/qewyxiww/project5/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метртов в одном километре
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
		log.Println(err)
		return ""
	}

	distanceM := float64(steps) * stepLength
	distanceKm := distanceM / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println(err)
		return ""
	}
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps, distanceKm, calories)

	return result
}
