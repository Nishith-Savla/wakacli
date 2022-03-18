package dto

func SumDuration(durationEntries ...DurationEntry) int {
	totalDuration := 0.0
	for _, durationEntry := range durationEntries {
		totalDuration += durationEntry.Duration
	}

	return int(totalDuration)
}
