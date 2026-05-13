package cli

import (
	"fmt"
	"time"

	"github.com/Judisk/daily-tracker-cli/internal/model"
	"github.com/Judisk/daily-tracker-cli/internal/validation"
)

type formField[T any] struct {
	prompt   string
	Validate func(string) (T, error)
}

type fields struct {
	wentToBed  formField[time.Time]
	fellAsleep formField[time.Time]
	wokeUp     formField[time.Time]
	tookMeds   formField[time.Time]

	sleepQuality formField[int]
	mood         formField[int]
	energy       formField[int]
	focus        formField[int]
	pills        formField[int]

	notes       formField[string]
	sideEffects formField[string]
}

func newFields() fields {

	return fields{

		wentToBed:  newTimeField("Went to bed"),
		fellAsleep: newTimeField("Fell asleep"),
		wokeUp:     newTimeField("Woke up"),
		tookMeds:   newTimeField("Took meds"),

		sleepQuality: newIntField("Sleep quality", model.MinValue, model.MaxValue),
		mood:         newIntField("Mood", model.MinValue, model.MaxValue),
		energy:       newIntField("Energy", model.MinValue, model.MaxValue),
		focus:        newIntField("Focus", model.MinValue, model.MaxValue),
		pills:        newIntField("Pills", model.MinValue, model.PillsMax),

		notes:       newStringField("Notes"),
		sideEffects: newStringField("Side Effects")}
}

func newIntField(prompt string, min, max int) formField[int] {
	field := formField[int]{
		prompt:   fmt.Sprintf("%s %d-%d", prompt, min, max),
		Validate: validation.Int(prompt, min, max),
	}
	return field
}

func newTimeField(prompt string) formField[time.Time] {
	field := formField[time.Time]{
		prompt:   prompt,
		Validate: validation.Time(),
	}
	return field
}
func newStringField(prompt string) formField[string] {
	field := formField[string]{
		prompt:   prompt,
		Validate: validation.String(),
	}
	return field
}
