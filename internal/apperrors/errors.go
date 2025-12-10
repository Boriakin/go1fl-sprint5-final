package apperrors

import "errors"

var (
	ErrInvalidSteps    = errors.New("steps must be greater than zero")
	ErrInvalidDuration = errors.New("duration must be greater than zero")
	ErrInvalidWeight   = errors.New("weight must be greater than zero")
	ErrInvalidHeight   = errors.New("height must be greater than zero")
	ErrZeroAvgSpeed    = errors.New("avgspeed is zero")
	ErrInvalidTraining = errors.New("неизвестный тип тренировки")
)
