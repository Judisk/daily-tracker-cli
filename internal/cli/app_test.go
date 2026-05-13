package cli

import (
	"bufio"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/Judisk/daily-tracker-cli/internal/model"
	"github.com/Judisk/daily-tracker-cli/internal/validation"
)

var testFieldInt = formField[int]{
	prompt:   "value",
	Validate: validation.Int("value", model.MinValue, model.MaxValue),
}
var testFieldTime = formField[time.Time]{
	prompt:   "value",
	Validate: validation.Time(),
}
var testFieldString = formField[string]{
	prompt:   "value",
	Validate: validation.String(),
}

func TestGetValueInt(t *testing.T) {
	test := "5"
	raw := strings.NewReader(test + "\n")
	reader := bufio.NewReader(raw)

	result, err := getValue(reader, testFieldInt)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
}

func TestGetValueIntRetry(t *testing.T) {
	raw := strings.NewReader("-5\nabc\n5\n")
	reader := bufio.NewReader(raw)
	result, err := getValue(reader, testFieldInt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
}

func TestGetValueEOF(t *testing.T) {
	raw := strings.NewReader("-5\nabc\n")
	reader := bufio.NewReader(raw)
	_, err := getValue(reader, testFieldInt)
	if err != io.EOF {
		t.Fatalf("expected EOF, got %v", err)
	}

}

func TestGetValueIntNTimes(t *testing.T) {

	raw := strings.NewReader(
		"abc\nss\n00:00\n1\n",
	)
	reader := bufio.NewReader(raw)

	result, err := getValue(reader, testFieldInt)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 1 {
		t.Errorf("expected 1, got %d", result)
	}
}

func TestGetValueTimeRetry(t *testing.T) {
	test := "12:30"
	raw := strings.NewReader("05:3\n221\n12.30\n" + test + "\n")
	reader := bufio.NewReader(raw)

	result, err := getValue(reader, testFieldTime)

	if err != nil {
		t.Fatalf("unexpected error")
	}
	want, _ := time.Parse("15:04", test)

	if !result.Equal(want) {
		t.Errorf("got %v, want %v", result, want)
	}
}

func TestGetString(t *testing.T) {
	test := "abc"
	raw := strings.NewReader(test + "\n")
	reader := bufio.NewReader(raw)

	result, err := getValue(reader, testFieldString)

	if err != nil {
		t.Fatalf("unexpected error")
	}
	if result != test {
		t.Fatalf("expected %q, got %q", test, result)
	}

}
