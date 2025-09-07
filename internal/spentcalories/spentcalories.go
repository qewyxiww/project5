package spentcalories

import (
	"fmt"
	"log"
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
	divorced := strings.Split(data, ",")
	if len(divorced) != 3 {
		return 0, "", 0, fmt.Errorf("неверные данные, длина слайса не равна 3")
	}
	firstElem := divorced[0]
	steps, err := strconv.Atoi(firstElem)
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка преобразования шагов")
	}
	secondElem := divorced[1]
	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("количество шагов должно быть больше нуля!")
	}
	thirdElem := divorced[2]
	duration, err := time.ParseDuration(thirdElem)
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка преобразования длительности")
	}
	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("продолжительность должна быть больше нуля!")
	}
	return steps, secondElem, duration, nil
}

func distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distanceM := float64(steps) * stepLength
	distanceKm := distanceM / mInKm
	return distanceKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	distanceKm := distance(steps, height)

	hours := duration.Hours()

	speed := distanceKm / hours

	return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	var distanceKm, speed, calories float64
	steps, activityType, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	switch activityType {
	case "Бег":
		distanceKm = distance(steps, height)
		speed = meanSpeed(steps, height, duration)
		calories, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
	case "Ходьба":
		distanceKm = distance(steps, height)
		speed = meanSpeed(steps, height, duration)
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", activityType)
	}

	durationHours := duration.Hours()
	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		activityType, durationHours, distanceKm, speed, calories)

	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть больше нуля!")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть больше нуля!")
	}
	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть больше нуля!")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть больше нуля!")
	}
	avgSpeed := meanSpeed(steps, height, duration)
	durationMin := duration.Minutes()
	calories := (weight * avgSpeed * durationMin) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть больше нуля!")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть больше нуля!")
	}
	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть больше нуля!")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть больше нуля!")
	}
	avgSpeed := meanSpeed(steps, height, duration)
	durationMin := duration.Minutes()
	beforeCalories := (weight * avgSpeed * durationMin) / minInH
	calories := beforeCalories * walkingCaloriesCoefficient
	return calories, nil
}
