package dto

type DurationEntry struct {
	Duration float64 `json:"duration"`
	Project  string  `json:"project"`
	Time     float64 `json:"time"`
}
