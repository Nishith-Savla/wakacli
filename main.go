package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const wakatimeAPIUrl = "https://wakatime.com/api/v1"

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	today := time.Now().Format("2006-01-02")
	durationsUrl, err := buildDurationsURL(os.Getenv("API_KEY"), today)

	response, err := http.Get(durationsUrl.String())
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	var decodedResponse map[string]interface{}

	err = json.NewDecoder(response.Body).Decode(&decodedResponse)
	if err != nil {
		panic(err)
	}

	totalDuration := getTotalDuration(decodedResponse)

	hourMinuteString := fmt.Sprintf("%d hr %d min", totalDuration/3600, totalDuration%3600/60)
	fmt.Println(hourMinuteString)
}

func getTotalDuration(decodedResponse map[string]interface{}) int {
	data := decodedResponse["data"]
	totalDuration := 0.0
	for _, durationEntry := range data.([]interface{}) {
		totalDuration += durationEntry.(map[string]interface{})["duration"].(float64)
	}

	return int(totalDuration)
}

func buildDurationsURL(apiKey string, date string) (*url.URL, error) {
	durationsUrl, err := url.Parse(fmt.Sprintf("%s/users/current/durations", wakatimeAPIUrl))
	if err != nil {
		return nil, err
	}

	// build query strings
	query := durationsUrl.Query()
	query.Add("api_key", apiKey)
	query.Add("date", date)
	durationsUrl.RawQuery = query.Encode()

	return durationsUrl, nil
}

func getPrettyResponse(m map[string]interface{}) (string, error) {
	prettyResponse, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "", err
	}
	return string(prettyResponse), nil
}
