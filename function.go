package function

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
    // This connects the name 'WorkoutAPI' to this function
	functions.HTTP("WorkoutAPI", WorkoutAPI)
}

// WorkoutAPI is the entry point
func WorkoutAPI(w http.ResponseWriter, r *http.Request) {
    // 1. Simple Auth Check
	clientKey := r.Header.Get("x-api-key")
	serverKey := os.Getenv("APP_SECRET_PASSWORD")

	if clientKey != serverKey {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

    // 2. Success Response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "authorized", 
		"message": "Hello from Go!",
	})
}
