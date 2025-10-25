package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

func main() {

	weather := getWeather(ApiKey, City)

	//your language
	//fmt.Printf("%s\n", string(makeWidget(&weather)))
	terminal(&weather)
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

	border := lipgloss.Border{
		Top:         "_/m\\_",
		Bottom:      "_/w\\_",
		Left:        "|(",
		Right:       "|)",
		TopLeft:     ".",
		TopRight:    ".",
		BottomLeft:  "\\",
		BottomRight: "/",
	}

	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#e1027e"))

	str := "\n Cidade: " + style.Render(weather.Location.Name) + " | "
	str += "Condição: " + style.Render(weather.Current.Condition.Text) + " \n"

	str += " Temperatura: " + style.Render(strconv.FormatFloat(weather.Current.TempC, 'f', 1, 64)) + "ºC | "
	str += "Sensação Térmica: " + style.Render(strconv.FormatFloat(weather.Current.FeelsLIkeC, 'f', 1, 64)) + "ºC \n"

	str += " Umidade: " + style.Render(strconv.Itoa(weather.Current.Humidity)) + "% | "
	str += "Precipitação: " + style.Render(strconv.FormatFloat(weather.Current.PrecipMM, 'f', 4, 64)) + "mm \n"

	fmt.Print(lipgloss.NewStyle().Border(border).BorderForeground(lipgloss.Color("#e1027e")).Render(str), "\n")

	//	fmt.Printf("Cidade: %s | Condição: %s\n", weather.Location.Name, weather.Current.Condition.Text)
	//	fmt.Printf("Temperatura: %.1fºC | Sensação Térmica: %.1fºC\n", weather.Current.TempC, weather.Current.FeelsLIkeC)
	//	fmt.Printf("Umidade: %d%% | Precipitação: %.3fmm\n", weather.Current.Humidity, weather.Current.PrecipMM)

}
