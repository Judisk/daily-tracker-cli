package sleep

import "time"

func Duration(wokeUp, fellAsleep time.Time) time.Duration {
	sleepDuration := wokeUp.Sub(fellAsleep)

	if sleepDuration < 0 {
		wokeUp = wokeUp.Add(24 * time.Hour)
		sleepDuration = wokeUp.Sub(fellAsleep)
	}

	return sleepDuration
}
