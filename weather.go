package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Weather struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
	Condition   string  `json:"condition"`
	Humidity    float64 `json:"humidity"`
	WindSpeed   float64 `json:"wind_speed"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func wetherhandler(w http.ResponseWriter, r *http.Request) {

	API_KEY := os.Getenv("API_KEY")
	w.Header().Set("Content-Type", "application/json")

	city := r.URL.Query().Get("city")
	if city == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "city parameter is required"})
		return
	}
	url := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/" + city + "?key=" + API_KEY

	resp, err := http.Get(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Failed to fetch weather data"})
	}
	defer resp.Body.Close()

	var data map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "failed to parse weather response"})

	}

	today := data["currentConditions"].(map[string]interface{})
	temp := today["temp"].(float64)
	condition := today["conditions"].(string)
	humidity := today["humidity"].(float64)
	windspeed := today["windspeed"].(float64)

	weather := Weather{
		City:        city,
		Temperature: temp,
		Condition:   condition,
		Humidity:    humidity,
		WindSpeed:   windspeed,
	}

	err = json.NewEncoder(w).Encode(weather)
	if err != nil {
		fmt.Println(err)
		return
	}
}
