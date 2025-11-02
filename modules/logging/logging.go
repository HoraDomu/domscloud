package logging

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const logFile = "./logs.txt"

func LogHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Use POST", http.StatusMethodNotAllowed)
		return
	}

	msg := r.URL.Query().Get("msg")
	if msg == "" {
		http.Error(w, "Missing msg param", http.StatusBadRequest)
		return
	}

	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Cannot write log", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	timestamp := time.Now().Format(time.RFC3339)
	fmt.Fprintf(f, "[%s] %s\n", timestamp, msg)

	fmt.Fprintf(w, "Logged: %s\n", msg)
}
