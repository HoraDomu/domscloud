package health

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"
)

type Health struct {
	CPU          string `json:"cpu"`
	Memory       string `json:"memory"`
	NumGoroutine int    `json:"goroutines"`
	Uptime       string `json:"uptime"`
}

var startTime = time.Now()

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	health := Health{
		CPU:          "N/A", // placeholder, can integrate gopsutil later
		Memory:       "N/A", // placeholder
		NumGoroutine: runtime.NumGoroutine(),
		Uptime:       time.Since(startTime).String(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}
