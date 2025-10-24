package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func main() {

	weather := getWeather(ApiKey, City)

	//your language
	fmt.Printf("%s", string(makeWidget(&weather)))
}

func makeWidget(weather *Weather) []byte {
	attributes := map[string]string{
		"icon": weather.Current.Condition.Icon,
		"temp": strconv.FormatFloat(weather.Current.TempC, 'f', 1, 64),
	}

	jsonObj, err := json.Marshal(attributes)

	if err != nil {
		log.Fatal(err)
	}

	return jsonObj
}

func terminal(weather *Weather) {
	fmt.Printf("Cidade: %s | Condição: %s\n", weather.Location.Name, weather.Current.Condition.Text)
	fmt.Printf("Temperatura: %.1fºC | Sensação Térmica: %.1fºC\n", weather.Current.TempC, weather.Current.FeelsLIkeC)
	fmt.Printf("Umidade: %d%% | Precipitação: %.3fmm\n", weather.Current.Humidity, weather.Current.PrecipMM)

}
