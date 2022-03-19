package dto

func GetDurationPerProject(durationEntries ...DurationEntry) map[string]float64 {
	durationPerProject := make(map[string]float64)
	for _, durationEntry := range durationEntries {
		durationPerProject[durationEntry.Project] += durationEntry.Duration
	}

	return durationPerProject
}

func GetTotalDuration(durationPerProject map[string]float64) int {
	totalDuration := 0.0
	for _, duration := range durationPerProject {
		totalDuration += duration
	}
	return int(totalDuration)
}
