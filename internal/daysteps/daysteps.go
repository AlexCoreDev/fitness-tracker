package daysteps

import (
	"fmt"
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

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	var lasting time.Duration
	s := strings.Split(data, ",")

	if len(s) < 2 {
		return 0, lasting, fmt.Errorf("invalid input: slice length is less 2")
	}

	countStep, err := strconv.Atoi(s[0])
	if err != nil {
		return 0, lasting, fmt.Errorf("invalid steps value: %w", err)
	}

	if countStep < 0 {
		return 0, lasting, fmt.Errorf("invalid steps value")
	}

	timeWalking, err := time.ParseDuration(s[1])
	if err != nil {
		return 0, lasting, fmt.Errorf("invalid duration value: %w", err)
	}

	return countStep, timeWalking, nil

}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию

	countStep, timeWalking, err := parsePackage(data)
	if err != nil {
		fmt.Println("parsePackage error:", err)
		return ""
	}

	if countStep < 0 {
		return ""
	}

	distance := stepLength * float64(countStep) / float64(mInKm)
	calories, err := spentcalories.WalkingSpentCalories(countStep, weight, height, timeWalking)
	if err != nil {
		return ""
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", countStep, distance, calories,
	)

}
