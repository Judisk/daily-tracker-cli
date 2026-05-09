package input

import (
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestInputInt(t *testing.T) {
	reader := strings.NewReader("5\n")

	result, err := Input(reader, parseAndValidateInt("mood", 0, 5))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}

}

func TestInputIntPills(t *testing.T) {
	test := 50
	reader := strings.NewReader(strconv.Itoa(test) + "\n")

	result, err := Input(reader, parseAndValidateInt("mood", 0, 50))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != test {
		t.Errorf("expected 5, got %d", result)
	}

}

func TestInputIntInvalidNegative(t *testing.T) {
	reader := strings.NewReader("-5\n")

	_, err := Input(reader, parseAndValidateInt("mood", 0, 5))

	if err == nil {
		t.Errorf("expected error , got nil")
	}

}

func TestInputIntInvalid(t *testing.T) {
	reader := strings.NewReader("abc\n")

	_, err := Input(reader, parseAndValidateInt("mood", 0, 5))

	if err == nil {
		t.Errorf("expected error , got nil")
	}
}

func TestInputTime(t *testing.T) {
	test := "00:00"
	reader := strings.NewReader(test + "\n")

	result, err := Input(reader, parseAndValidateTime())
	if err != nil {
		t.Fatalf("unexpected error")
	}
	want, _ := time.Parse("15:04", test)

	if !result.Equal(want) {
		t.Errorf("got %v, want %v", result, want)
	}
}

func TestInputInvalidTime(t *testing.T) {
	test := "abc"
	reader := strings.NewReader(test + "\n")

	_, err := Input(reader, parseAndValidateTime())
	if err == nil {
		t.Errorf("expected error")
	}

}

func TestInputString(t *testing.T) {
	test := "abc"
	reader := strings.NewReader(test + "\n")

	result, err := Input(reader, stringValidation())
	if err != nil {
		t.Errorf("unexpected error")
	}
	if result != test {
		t.Errorf("expected %s, got %s", test, result)
	}

}
