package input

import (
	"bufio"
	"strings"
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
