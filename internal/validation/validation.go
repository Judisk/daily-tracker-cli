package validation

import (
	"fmt"
	"strconv"
	"time"
)

func Int(fieldName string, min, max int) func(string) (int, error) {
	return func(s string) (int, error) {
		num, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		if num < min || num > max {
			return 0, fmt.Errorf("%s must be %d-%d", fieldName, min, max)
		}

		return num, nil
	}
}

func Time() func(string) (time.Time, error) {
	return func(s string) (time.Time, error) {
		return time.Parse("15:04", s)
	}
}

func String() func(string) (string, error) {
	return func(s string) (string, error) {
		return s, nil
	}
}
