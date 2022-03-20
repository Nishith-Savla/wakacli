package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/Nishith-Savla/wakacli/common"
	"github.com/Nishith-Savla/wakacli/dto"
)

func SetAPIKey(apiKey string) error {
	dirname, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	data, err := json.MarshalIndent(map[string]string{"apikey": apiKey}, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = os.WriteFile(filepath.Join(dirname, ".wakacli.json"), data, 0600)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func GetAPIKey() (string, error) {
	var err error
	var dirname string

	if dirname, err = os.UserHomeDir(); err != nil {
		return "", err
	}

	var data []byte
	if data, err = os.ReadFile(filepath.Join(dirname, ".wakacli.json")); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File not found")
			return "", errors.New("please first set the API Key using `wakacli apikey <your-wakatime-api-key>")
		}
		fmt.Println(err)
		return "", err
	}

	var unmarshalledData struct {
		ApiKey string `json:"apikey"`
	}
	if err = json.Unmarshal(data, &unmarshalledData); err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return unmarshalledData.ApiKey, nil
}

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

func secToHrMinSec(seconds int) (int, int, int) {
	return seconds / 3600, seconds % 3600 / 60, seconds % 60
}

func getDuration(date string, includeSeconds bool) (string, error) {
	apiKey, err := GetAPIKey()
	if err != nil {
		return "", err
	}

	durationsUrl, err := buildURL(
		apiKey,
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

	var output strings.Builder
	durationPerProject := dto.GetDurationPerProject(durationEntries...)

	var maxLength int
	for project := range durationPerProject {
		length := len(project)
		if length > maxLength {
			maxLength = length
		}
	}

	for project, duration := range durationPerProject {
		hr, min, _ := secToHrMinSec(int(duration))
		output.WriteString(fmt.Sprintf("%-*s | %02d hr %02d min\n", maxLength, project, hr, min))
	}
	output.WriteString("\n")

	totalDuration := dto.GetTotalDuration(durationPerProject)

	hr, min, sec := secToHrMinSec(totalDuration)
	readableDuration := fmt.Sprintf("%d hr %d min", hr, min)
	if includeSeconds {
		readableDuration += fmt.Sprintf(" %d sec", sec)
	}

	output.WriteString(readableDuration)

	return output.String(), nil
}
