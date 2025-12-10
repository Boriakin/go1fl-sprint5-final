package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/apperrors"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(data string) (err error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return fmt.Errorf("invalid data: %s", data)
	}

	ds.Steps, err = strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	if ds.Steps <= 0 {
		return apperrors.ErrInvalidSteps
	}

	ds.Duration, err = time.ParseDuration(parts[1])
	if err != nil {
		return err
	}
	if ds.Duration <= 0 {
		return apperrors.ErrInvalidDuration
	}

	return nil
}

func (ds *DaySteps) ActionInfo() (string, error) {
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	dist := spentenergy.Distance(ds.Steps, ds.Height)

	info := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, dist, calories)

	return info, nil
}
