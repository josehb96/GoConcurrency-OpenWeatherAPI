package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const apiKey = "..."

func fetchWeather(city string, ch chan<- string, wg *sync.WaitGroup) interface{} {

	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	defer wg.Done() // This ensures us that will close and send the signal to our primitive which is our wait group to finish off and that this go routine has completed

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching weather for %s: %s\n", city, err)
		return data
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("Error decoding weather data for %s: %s\n", city, err)
		return data
	}

	ch <- fmt.Sprintf("This is the %s", city)

	return data

}

func main() {

	startNow := time.Now()

	cities := []string{"Chicago", "Rome", "Miami", "New York"}

	ch := make(chan string)
	var wg sync.WaitGroup

	for _, city := range cities {
		wg.Add(1) // Tell our primitive that we are adding a go routine
		go fetchWeather(city, ch, &wg)
	}

	// Declare another go routine to tell our primitive to make sure to wait for all of our go routines to finish
	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}

	fmt.Println("This operation took:", time.Since(startNow))

}
