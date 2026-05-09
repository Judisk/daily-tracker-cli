package input

import (
	"bufio"
	"io"
	"strings"
)

/*
var filedNameInt = map[string]int{"mood": 1, "sleep quality": 1, "energy": 1, "focus": 1}

const fieldNamePills = "pills"

var fieldNameTime = map[string]int{"went to bed": 1, "took meds": 1, "fellAsleep": 1, "woke up": 1}

	func InputInt(r io.Reader, fieldName string, min, max int) (int, error) {
		reader := bufio.NewReader(r)

		s, err := reader.ReadString('\n')
		if err != nil {
			return 0, err
		}
		s = strings.TrimSpace(s)

		value, err := parseAndValidateInt(s, fieldName, min, max)
		if err != nil {
			return 0, err
		}

		return value, nil
	}

	func InputTime(r io.Reader) (t time.Time, err error) {
		reader := bufio.NewReader(r)
		s, err := reader.ReadString('\n')
		if err != nil {
			return t, err
		}
		s = strings.TrimSpace(s)
		return parseAndValidateTime(s)
	}
*/

func Input[T any](r io.Reader, f func(string) (T, error)) (T, error) {
	reader := bufio.NewReader(r)
	var zero T
	s, err := reader.ReadString('\n')
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
