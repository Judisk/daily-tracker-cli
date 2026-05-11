package model

import "time"

type Record struct {
	//время
	Date          time.Time     `json:"date"`
	WentToBed     time.Time     `json:"went_to_bed"`
	FellAsleep    time.Time     `json:"fell_asleep"`
	WokeUp        time.Time     `json:"woke_up"`
	SleepDuration time.Duration `json:"sleep_duration"`
	//оценки
	SleepQuality int `json:"sleep_quality"`
	Mood         int `json:"mood"`
	Energy       int `json:"energy"`
	Focus        int `json:"focus"`
	//сколько осталось таблеток
	Pills    int       `json:"pills"`
	TookMeds time.Time `json:"took_meds"`
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
