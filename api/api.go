package api

import "time"

func GetToday(includeSeconds bool) (string, error) {
	today := time.Now().Format("2006-01-02")
	return getDuration(today, includeSeconds)
}

func GetYesterday(includeSeconds bool) (string, error) {
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	return getDuration(yesterday, includeSeconds)
}
