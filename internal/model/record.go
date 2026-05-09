package model

type Record struct {
	//время
	Date       string `json:"date"`
	WentToBed  string `json:"went_to_bed"`
	TookMeds   string `json:"took_meds"`
	FellAsleep string `json:"fell_asleep"`
	WokeUp     string `json:"woke_up"`
	//оценки
	SleepQuality int `json:"sleep_quality"`
	Mood         int `json:"mood"`
	Energy       int `json:"energy"`
	Focus        int `json:"focus"`
	//сколько осталось таблеток
	Pills int `json:"pills"`
	//записи
	Notes       string `json:"notes"`
	SideEffects string `json:"side_effects"`
}

const (
	MinValue = 0
	MaxValue = 5
)
const (
	PillsMax          = 50
	PillsLowThreshold = 7
)
