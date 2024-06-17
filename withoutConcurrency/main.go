package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// apiKey is a constant for the OpenWeatherMap API key
const apiKey = "..."

// fetchWeather fetches the weather data for a given city
func fetchWeather(city string) interface{} {

	// Define a struct to hold the weather data
	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	// Construct the URL for the API request
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	// Send a GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		// Handle any errors that occur during the request
		fmt.Printf("Error fetching weather for %s: %s\n", city, err)
		return data
	}

	// Close the response body when we're done with it
	defer resp.Body.Close()

	// Decode the JSON response into the data struct
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		// Handle any errors that occur during JSON decoding
		fmt.Printf("Error decoding weather data for %s: %s\n", city, err)
		return data
	}

	// Return the weather data
	return data

}

// main is the entry point for the program
func main() {

	// Record the start time for the operation
	startNow := time.Now()

	// Define a list of cities to fetch weather data for
	cities := []string{"Madrid", "Rome", "Miami", "New York"}

	// Iterate over each city and fetch its weather data
	for _, city := range cities {
		data := fetchWeather(city)
		fmt.Println("This is the data", data)
	}

	// Print the time taken for the operation
	fmt.Println("This operation took:", time.Since(startNow))

}
