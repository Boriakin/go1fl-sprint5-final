package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/apperrors"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(data string) (err error) {
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return fmt.Errorf("invalid data: %s", data)
	}

	t.Steps, err = strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	if t.Steps <= 0 {
		return apperrors.ErrInvalidSteps
	}

	t.TrainingType = parts[1]

	t.Duration, err = time.ParseDuration(parts[2])
	if err != nil {
		return err
	}
	if t.Duration <= 0 {
		return apperrors.ErrInvalidDuration
	}

	return nil
}

func (t *Training) ActionInfo() (string, error) {
	calories, err := t.calcCalories()
	if err != nil {
		return "", err
	}

	dist := spentenergy.Distance(t.Steps, t.Height)

	avgSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	if avgSpeed == 0 {
		return "", apperrors.ErrZeroAvgSpeed
	}

	info := fmt.Sprintf("Тип тренировки: %s\n"+
		"Длительность: %.2f ч.\n"+
		"Дистанция: %.2f км.\n"+
		"Скорость: %.2f км/ч\n"+
		"Сожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), dist, avgSpeed, calories)

	return info, nil
}

func (t *Training) calcCalories() (float64, error) {
	switch t.TrainingType {
	case "Бег":
		return spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

	case "Ходьба":
		return spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

	default:
		return 0, apperrors.ErrInvalidTraining
	}
}
