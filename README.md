# WeatherFetcher
A simple command-line tool written in Go that fetches current weather data from the [OpenWeatherMap API](https://openweathermap.org/)

# Features
* Fetches weather by city name (e.g., Dehradun,IN)
* Falls back to interactive input if no city is provided
* Displays:
>* Location (city, country, coordinates)
>* Temperature (current, feels like, min, max)
>* Humidity and pressure
>* Wind speed & direction
>* Cloud coverage
>* Sunrise & sunset times
>* Weather condition description

# Installation
Clone the repo:
```bash
git clone https://github.com/Probatio-Diabolica/WeatherFetcher.git
cd WeatherFetcher
```
Install dependencies:
```bash
go mod tidy
```

# Setup
1. Get a free API key from [OpenWeatherMap](https://openweathermap.org/)
2. Create a .env file in the project root:
```env
API_KEYS=your_api_key_here
```

# Usage

Run with a city flag:
```bash
go run . -city="city_mame,country_Code"
```

Or let it ask you:

```bash
go run .
Enter your city
```

# Example output

```yaml
Location: Dehradun, IN (Lat: 30.33, Lon: 78.04)
Temperature: 23.3°C (Feels like 24.1°C)
Min/Max: 22.5°C / 24.0°C
Humidity: 83%
Pressure: 1008hPa
Wind: 1.9 m/s at 61°
Cloud Cover: 72%
Sunrise: 05:56
Sunset: 18:07
Condition: broken clouds
```

# Notes
>* Uses [godotenv](https://github.com/joho/godotenv) for .env loading
>* By default, temperatures are returned in Kelvin by the API and converted to Celsius in code. You can also add &units=metric in the API URL to get Celsius directly.
