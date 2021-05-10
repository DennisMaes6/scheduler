package schedule_generator

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jorensjongers/scheduler/backend/model"
	"github.com/pkg/errors"
)

func integerArrayToString(array []int32) string {
	result := ""
	for _, h := range array {
		result += fmt.Sprintf("%s,", strconv.Itoa(int(h)))
	}
	return result
}

func stringToIntArray(str string) ([]int32, error) {
	if str == "" {
		return []int32{}, nil
	}
	splitStr := strings.Split(strings.TrimSuffix(str, ","), ",")
	result := []int32{}
	for _, s := range splitStr {
		i, err := strconv.Atoi(s)
		if err != nil {
			return []int32{}, err
		}
		result = append(result, int32(i))
	}

	return result, nil
}

func dateToString(date model.DayDate) string {
	return fmt.Sprintf("%d-%d-%d", date.Day, date.Month, date.Year)
}

func parseDate(rawDate string) (model.DayDate, error) {
	ints := strings.Split(rawDate, "-")
	day, err := strconv.Atoi(ints[0])
	if err != nil {
		return model.DayDate{}, errors.Wrap(err, fmt.Sprintf("failed persing date day: %s", ints[0]))
	}
	month, err := strconv.Atoi(ints[1])
	if err != nil {
		return model.DayDate{}, errors.Wrap(err, fmt.Sprintf("failed persing date month: %s", ints[1]))
	}
	year, err := strconv.Atoi(ints[2])
	if err != nil {
		return model.DayDate{}, errors.Wrap(err, fmt.Sprintf("failed persing date year: %s", ints[2]))
	}
	return model.DayDate{
		Day:   int32(day),
		Month: int32(month),
		Year:  int32(year),
	}, nil

}
