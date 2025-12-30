package function

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("WorkoutAPI", WorkoutAPI)
}

// setCORSHeaders sets the CORS headers for all responses
func setCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, x-api-key")
}

// WorkoutAPI is the entry point for the Cloud Function
func WorkoutAPI(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)

	// Handle preflight requests
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Auth check
	clientKey := r.Header.Get("x-api-key")
	serverKey := os.Getenv("APP_SECRET_PASSWORD")

	if clientKey != serverKey {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get Firestore client
	ctx := context.Background()
	projectID := os.Getenv("GCP_PROJECT_ID")
	if projectID == "" {
		projectID = os.Getenv("GOOGLE_CLOUD_PROJECT") // Fallback for Cloud Functions
	}

	client, err := GetFirestoreClient(ctx, projectID)
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}

	// Route requests
	path := r.URL.Path
	method := r.Method

	// /indoor_sessions routes
	if strings.HasPrefix(path, "/indoor_sessions") {
		sessionID := ParseSessionID(path)

		switch {
		case method == "GET" && sessionID == "":
			ListIndoorSessions(w, r, client)
		case method == "GET" && sessionID != "":
			GetIndoorSession(w, r, client, sessionID)
		case method == "POST" && sessionID == "":
			CreateIndoorSession(w, r, client)
		case method == "PUT" && sessionID != "":
			UpdateIndoorSession(w, r, client, sessionID)
		case method == "DELETE" && sessionID != "":
			DeleteIndoorSession(w, r, client, sessionID)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}

	// /outdoor_sessions routes
	if strings.HasPrefix(path, "/outdoor_sessions") {
		sessionID := ParseOutdoorSessionID(path)

		switch {
		case method == "GET" && sessionID == "":
			ListOutdoorSessions(w, r, client)
		case method == "GET" && sessionID != "":
			GetOutdoorSession(w, r, client, sessionID)
		case method == "POST" && sessionID == "":
			CreateOutdoorSession(w, r, client)
		case method == "PUT" && sessionID != "":
			UpdateOutdoorSession(w, r, client, sessionID)
		case method == "DELETE" && sessionID != "":
			DeleteOutdoorSession(w, r, client, sessionID)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}

	// /fingerboard_sessions routes
	if strings.HasPrefix(path, "/fingerboard_sessions") {
		sessionID := ParseFingerboardSessionID(path)

		switch {
		case method == "GET" && sessionID == "":
			ListFingerboardSessions(w, r, client)
		case method == "GET" && sessionID != "":
			GetFingerboardSession(w, r, client, sessionID)
		case method == "POST" && sessionID == "":
			CreateFingerboardSession(w, r, client)
		case method == "PUT" && sessionID != "":
			UpdateFingerboardSession(w, r, client, sessionID)
		case method == "DELETE" && sessionID != "":
			DeleteFingerboardSession(w, r, client, sessionID)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}

	// /competition_sessions routes
	if strings.HasPrefix(path, "/competition_sessions") {
		sessionID := ParseCompetitionSessionID(path)

		switch {
		case method == "GET" && sessionID == "":
			ListCompetitionSessions(w, r, client)
		case method == "GET" && sessionID != "":
			GetCompetitionSession(w, r, client, sessionID)
		case method == "POST" && sessionID == "":
			CreateCompetitionSession(w, r, client)
		case method == "PUT" && sessionID != "":
			UpdateCompetitionSession(w, r, client, sessionID)
		case method == "DELETE" && sessionID != "":
			DeleteCompetitionSession(w, r, client, sessionID)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}

	// Default: not found
	http.Error(w, "Not found", http.StatusNotFound)
}
