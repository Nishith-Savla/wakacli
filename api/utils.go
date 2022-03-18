package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/Nishith-Savla/wakacli/common"
	"github.com/Nishith-Savla/wakacli/dto"
)

func buildURL(apiKey string, predicate string, queryParams map[string]string) (*url.URL, error) {
	durationsUrl, err := url.Parse(fmt.Sprintf("%s/users/current/%s", common.WakaTimeAPIUrl, predicate))
	if err != nil {
		return nil, err
	}

	// build query strings
	query := durationsUrl.Query()
	query.Add("api_key", apiKey)
	for key, value := range queryParams {
		query.Add(key, value)
	}
	durationsUrl.RawQuery = query.Encode()

	return durationsUrl, nil
}

func getDuration(date string, includeSeconds bool) (string, error) {
	durationsUrl, err := buildURL(
		os.Getenv("API_KEY"),
		"durations",
		map[string]string{"date": date},
	)

	var response *http.Response
	if response, err = http.Get(durationsUrl.String()); err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	var decodedResponse map[string]interface{}

	if err = json.NewDecoder(response.Body).Decode(&decodedResponse); err != nil {
		return "", err
	}

	var data []byte
	// Convert slice to json
	if data, err = json.Marshal(decodedResponse["data"]); err != nil {
		return "", err
	}

	var durationEntries []dto.DurationEntry
	// Convert json to struct
	if err = json.Unmarshal(data, &durationEntries); err != nil {
		return "", err
	}

	totalDuration := dto.SumDuration(durationEntries...)
	readableDuration := fmt.Sprintf("%d hr %d min", totalDuration/3600, totalDuration%3600/60)
	if includeSeconds {
		readableDuration += fmt.Sprintf(" %d sec", totalDuration%60)
	}
	return readableDuration, nil
}
