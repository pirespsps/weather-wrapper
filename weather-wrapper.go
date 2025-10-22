package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Weather struct {
	Location struct {
		Name string `json:"name"`
	}
	Current struct {
		Condition struct {
			Text string `json:"text"`
		}

		TempC      float64 `json:"temp_c"`
		FeelsLIkeC float64 `json:"feelslike_c"`
		Humidity   int     `json:"humidity"`
		PrecipMM   float64 `json:"precip_mm"`
	}
}

func getWeather(key string, city string) Weather {
	var url = "http://api.weatherapi.com/v1/current.json?key=" + key + "&q=" + city + "&aqi=no"
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var weather Weather
	json.Unmarshal(data, &weather)

	return weather
}
