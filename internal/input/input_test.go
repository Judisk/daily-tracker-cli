package input

import (
	"bufio"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/Judisk/daily-tracker-cli/internal/validation"
)

func TestInputInt(t *testing.T) {
	raw := strings.NewReader("5\n")
	reader := bufio.NewReader(raw)
	result, err := Input(reader, validation.Int("mood", 0, 5))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}

}

func TestInputIntPills(t *testing.T) {
	test := 50
	raw := strings.NewReader(strconv.Itoa(test) + "\n")
	reader := bufio.NewReader(raw)
	result, err := Input(reader, validation.Int("pills", 0, 50))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != test {
		t.Errorf("expected 5, got %d", result)
	}

}

func TestInputIntInvalidNegative(t *testing.T) {
	raw := strings.NewReader("-5\n")
	reader := bufio.NewReader(raw)
	_, err := Input(reader, validation.Int("mood", 0, 5))

	if err == nil {
		t.Errorf("expected error , got nil")
	}

}

func TestInputIntInvalid(t *testing.T) {
	raw := strings.NewReader("abc\n")
	reader := bufio.NewReader(raw)

	_, err := Input(reader, validation.Int("mood", 0, 5))

	if err == nil {
		t.Errorf("expected error , got nil")
	}
}

func TestInputTime(t *testing.T) {
	test := "00:00"
	raw := strings.NewReader(test + "\n")
	reader := bufio.NewReader(raw)

	result, err := Input(reader, validation.Time())
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
	raw := strings.NewReader(test + "\n")
	reader := bufio.NewReader(raw)

	_, err := Input(reader, validation.Time())
	if err == nil {
		t.Errorf("expected error")
	}

}

func TestInputString(t *testing.T) {
	test := "abc"
	raw := strings.NewReader(test + "\n")
	reader := bufio.NewReader(raw)

	result, err := Input(reader, validation.String())
	if err != nil {
		t.Errorf("unexpected error")
	}
	if result != test {
		t.Errorf("expected %s, got %s", test, result)
	}

}
