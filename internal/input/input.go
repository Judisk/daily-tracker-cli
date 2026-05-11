package input

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Input[T any](r *bufio.Reader, f func(string) (T, error)) (T, error) {
	var zero T

	s, err := r.ReadString('\n')
	if err != nil {
		return zero, err
	}
	s = strings.TrimSpace(s)

	result, err := f(s)
	if err != nil {
		return zero, err
	}
	return result, nil

}

func IntValidator(fieldName string, min, max int) func(string) (int, error) {
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

func TimeValidator() func(string) (time.Time, error) {
	return func(s string) (time.Time, error) {
		return time.Parse("15:04", s)
	}
}

func StringValidation() func(string) (string, error) {
	return func(s string) (string, error) {
		return s, nil
	}
}
