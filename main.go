package main

import (
	"fmt"
)

func main() {

	weather := getWeather(ApiKey, City)

	//your language
	fmt.Printf("Cidade: %s | Condição: %s\n", weather.Location.Name, weather.Current.Condition.Text)
	fmt.Printf("Temperatura: %.1fºC | Sensação Térmica: %.1fºC\n", weather.Current.TempC, weather.Current.FeelsLIkeC)
	fmt.Printf("Umidade: %d%% | Precipitação: %fmm\n", weather.Current.Humidity, weather.Current.PrecipMM)

}
