package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type DiffTime struct {
	DayOfWeek   string `json:"day_of_week"`
	DayOfMonth  int    `json:"day_of_month"`
	Month       string `json:"month"`
	Hour        int    `json:"hour"`
	Year        int    `json:"year"`
	Second      int    `json:"second"`
	Minute      int    `json:"minute"`
}

func getCurrentTime() DiffTime {
	n := time.Now()
	return DiffTime{
		DayOfWeek:   n.Weekday().String(),
		DayOfMonth:  n.Day(),
		Month:       n.Month().String(),
		Year:        n.Year(),
		Hour:        n.Hour(),
		Minute:      n.Minute(),
		Second:      n.Second(),
	}
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	current := getCurrentTime()
	fmt.Println(current)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(current)
}

func main() {
	http.HandleFunc("/time/RFC3339", timeHandler)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":80", nil)
}
