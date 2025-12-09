package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

var (
	ErrInvalidSteps    = errors.New("steps must be greater than zero")
	ErrInvalidDuration = errors.New("duration must be greater than zero")
	ErrInvalidWeight   = errors.New("weight must be greater than zero")
	ErrInvalidHeight   = errors.New("height must be greater than zero")
	ErrZeroAvgSpeed    = errors.New("avgspeed is zero")
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, ErrInvalidSteps
	}

	if weight <= 0 {
		return 0, ErrInvalidWeight
	}

	if height <= 0 {
		return 0, ErrInvalidHeight
	}

	if duration <= 0 {
		return 0, ErrInvalidDuration
	}

	avgSpeed := MeanSpeed(steps, height, duration)
	if avgSpeed == 0 {
		return 0, ErrZeroAvgSpeed
	}

	durationInMin := duration.Minutes()

	calories := weight * avgSpeed * durationInMin / minInH

	return calories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, ErrInvalidSteps
	}

	if weight <= 0 {
		return 0, ErrInvalidWeight
	}

	if height <= 0 {
		return 0, ErrInvalidHeight
	}

	if duration <= 0 {
		return 0, ErrInvalidDuration
	}

	avgSpeed := MeanSpeed(steps, height, duration)
	if avgSpeed == 0 {
		return 0, ErrZeroAvgSpeed
	}

	durationInMin := duration.Minutes()

	calories := weight * avgSpeed * durationInMin / minInH

	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}

	stepLength := height * stepLengthCoefficient

	return float64(steps) * stepLength / mInKm
}
