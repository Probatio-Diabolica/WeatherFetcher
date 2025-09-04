package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type WeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
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

	var w WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&w); err != nil {
		log.Fatal("Failed to decode JSON: ", err)
	}

	fmt.Printf("Weather in %s: %.1f°C - %s\n",
		w.Name,
		w.Main.Temp-273.15,
		w.Weather[0].Description,
	)
}
