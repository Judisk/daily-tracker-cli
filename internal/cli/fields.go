package cli

import (
	"fmt"
	"time"

	"github.com/Judisk/daily-tracker-cli/internal/model"
	"github.com/Judisk/daily-tracker-cli/internal/validation"
)

type Field[T any] struct {
	prompt   string
	Validate func(string) (T, error)
}

type Fields struct {
	WentToBed  Field[time.Time]
	FellAsleep Field[time.Time]
	WokeUp     Field[time.Time]
	TookMeds   Field[time.Time]

	SleepQuality Field[int]
	Mood         Field[int]
	Energy       Field[int]
	Focus        Field[int]
	Pills        Field[int]

	Notes       Field[string]
	SideEffects Field[string]
}

func NewFields() Fields {

	return Fields{

		WentToBed:  newTimeField("Went to bed"),
		FellAsleep: newTimeField("Fell asleep"),
		WokeUp:     newTimeField("Woke up"),
		TookMeds:   newTimeField("Took meds"),

		SleepQuality: newIntField("Sleep quality", model.MinValue, model.MaxValue),
		Mood:         newIntField("Mood", model.MinValue, model.MaxValue),
		Energy:       newIntField("Energy", model.MinValue, model.MaxValue),
		Focus:        newIntField("Focus", model.MinValue, model.MaxValue),
		Pills:        newIntField("Pills", model.MinValue, model.PillsMax),

		Notes:       newStringField("Notes"),
		SideEffects: newStringField("Side Effects")}
}

func newIntField(prompt string, min, max int) Field[int] {
	field := Field[int]{
		prompt:   fmt.Sprintf("%s %d-%d", prompt, min, max),
		Validate: validation.Int(prompt, min, max),
	}

	return field
}

func newTimeField(prompt string) Field[time.Time] {
	field := Field[time.Time]{
		prompt:   prompt,
		Validate: validation.Time(),
	}

	return field
}
func newStringField(prompt string) Field[string] {
	field := Field[string]{
		prompt:   prompt,
		Validate: validation.String(),
	}

	return field
}
