package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/Judisk/daily-tracker-cli/internal/input"
	"github.com/Judisk/daily-tracker-cli/internal/model"
)

func TestGetValueInt1Time(t *testing.T) {
	test := "5"
	raw := strings.NewReader(test + "\n")
	reader := bufio.NewReader(raw)

	result, err := getValue(reader, "Mood 0-5", input.ParseAndValidateInt("mood", model.MinValue, model.MaxValue))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
}

func TestInvalidGetValueIntNegative(t *testing.T) {
	test := "-5"
	raw := strings.NewReader(test + "\n")
	reader := bufio.NewReader(raw)

	_, err := getValue(reader, "Mood 0-5", input.ParseAndValidateInt("mood", model.MinValue, model.MaxValue))

	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestGetValueIntNTimes(t *testing.T) {

	raw := strings.NewReader(
		"abc\nss\n00:00\n1\n",
	)
	reader := bufio.NewReader(raw)

	result, err := getValue(
		reader,
		"Mood 0-5",
		input.ParseAndValidateInt("mood",
			model.MinValue,
			model.MaxValue),
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 1 {
		t.Errorf("expected 1, got %d", result)
	}
}

func TestInvalidGetValueTime1Time(t *testing.T) {
	test := "05:3"
	raw := strings.NewReader(test + "\n")
	reader := bufio.NewReader(raw)

	_, err := getValue(reader, "Mood 0-5", input.ParseAndValidateTime())

	if err == nil {
		t.Errorf("expected error")
	}
}

func TestGetString(t *testing.T) {
	test := "abc"
	raw := strings.NewReader(test + "\n")
	reader := bufio.NewReader(raw)

	result, err := getValue(reader, "Mood 0-5", input.StringValidation())

	if err != nil {
		t.Fatalf("unexpected error")
	}
	if result != test {
		t.Fatalf("unexpected error")
	}

}
