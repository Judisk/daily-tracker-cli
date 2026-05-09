package input

import (
	"fmt"
	"strconv"
	"time"
)

/*
func (r Record) Validate() error {

		if err := ValidateRange(r.SleepQuality, MinValue, MaxValue, "sleep quality"); err != nil {
			return err
		}
		if err := ValidateRange(r.Mood, MinValue, MaxValue, "mood"); err != nil {
			return err
		}
		if err := ValidateRange(r.Energy, MinValue, MaxValue, "energy"); err != nil {
			return err
		}
		if err := ValidateRange(r.Focus, MinValue, MaxValue, "focus"); err != nil {
			return err
		}
		if err := ValidateRange(r.Pills, PillsMin, PillsMax, "pills"); err != nil {
			return err
		}
		return nil
	}

	func ValidateRange(value, min, max int, fieldName string) error {
		if value < min || value > max {
			return fmt.Errorf("%s must be %d-%d", fieldName, min, max)
		}
		return nil
	}
*/
func parseAndValidateInt(fieldName string, min, max int) func(string) (int, error) {
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

func parseAndValidateTime() func(string) (time.Time, error) {
	return func(s string) (time.Time, error) {
		return time.Parse("15:04", s)
	}
}

func stringValidation() func(string) (string, error) {
	return func(s string) (string, error) {
		return s, nil
	}
}
