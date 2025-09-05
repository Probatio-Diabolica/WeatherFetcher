package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type WeatherResponse struct {
	Name string `json:"name"`
	Sys  struct {
		Country string `json:"country"`
		Sunrise int64  `json:"sunrise"`
		Sunset  int64  `json:"sunset"`
	} `json:"sys"`
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"sunset"`
	}
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func main() {
	godotenv.Load()
	apikey := os.Getenv("API_KEY")
	if apikey == "" {
		log.Fatal("Missing API_KEY in .env")
	}

	city := flag.String("city", "", "city name")
	flag.Parse()

	if *city == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your city\n")
		ci, _ := reader.ReadString('\n')
		*city = ci
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", *city, apikey)

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("Failed to fetch data: ", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("API request failed with staus; %s", resp.Status)
	}

	var weather WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		log.Fatal("Failed to decode JSON: ", err)
	}

	fmt.Printf(
		"Location: %s, %s (Lat: %.2f, Lon: %.2f)\n"+
			"Temperature: %.1f°C (Feels like %.1f°C)\n"+
			"Min/Max: %.1f°C / %.1f°C\n"+
			"Humidity: %d%%\n"+
			"Pressure: %dhPa\n"+
			"Wind: %.1f m/s at %d°\n"+
			"Cloud Cover: %d%%\n"+
			"Sunrise: %s\n"+
			"Sunset: %s\n"+
			"Condition: %s\n",
		weather.Name, weather.Sys.Country,
		weather.Coord.Lat, weather.Coord.Lon,
		weather.Main.Temp-273.15, weather.Main.FeelsLike-273.15,
		weather.Main.TempMin-273.15, weather.Main.TempMax-273.15,
		weather.Main.Humidity,
		weather.Main.Pressure,
		weather.Wind.Speed, weather.Wind.Deg,
		weather.Clouds.All,
		time.Unix(weather.Sys.Sunrise, 0).Format("15:04"),
		time.Unix(weather.Sys.Sunset, 0).Format("15:04"),
		weather.Weather[0].Description,
	)

}
