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
	// TODO: реализовать функцию
	s := strings.Split(data, ",")
	var lasting time.Duration
	if len(s) < 3 {
		return 0, "", lasting, fmt.Errorf("invalid input: slice length in less 3")
	}

	countStep, err := strconv.Atoi(s[0])
	if err != nil {
		return 0, "", lasting, fmt.Errorf("invalid steps value: %w", err)
	}

	typeActivity := s[1]

	timeWalking, err := time.ParseDuration(s[2])
	if err != nil {
		return 0, "", lasting, fmt.Errorf("invalid duration value: %w", err)
	}

	return countStep, typeActivity, timeWalking, nil

}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	lenghtStep := height * stepLengthCoefficient

	return float64(steps) * lenghtStep / float64(mInKm)
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 || steps <= 0 {
		return 0
	}

	return distance(steps, height) / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	countStep, typeActivity, timeWalking, err := parseTraining(data)
	if err != nil {
		log.Println("parseTraining error:", err)
		return "", fmt.Errorf("TrainingInfo: %w", err)
	}

	dist := distance(countStep, height)
	averageSpeed := meanSpeed(countStep, height, timeWalking)
	var calories float64

	switch typeActivity {
	case "Бег":
		calories, err = RunningSpentCalories(countStep, weight, height, timeWalking)
		if err != nil {
			return "", fmt.Errorf("failed to calculate running calories: %w", err)
		}

	case "Ходьба":
		calories, err = WalkingSpentCalories(countStep, weight, height, timeWalking)
		if err != nil {
			return "", fmt.Errorf("filed to calculate walking calories: %w", err)
		}

	default:
		return "", fmt.Errorf("unknown type of training")

	}

	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", typeActivity, timeWalking.Hours(), dist, averageSpeed, calories,
	), nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps < 0 {
		return 0, fmt.Errorf("invalid steps value")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("invalid duration value")
	}

	averageSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	return (weight * averageSpeed * durationInMinutes) / float64(minInH), nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps < 0 {
		return 0, fmt.Errorf("invalid steps value")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("invalid duration value")
	}

	averageSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * durationInMinutes * averageSpeed) / float64(minInH)

	return calories * walkingCaloriesCoefficient, nil

}
