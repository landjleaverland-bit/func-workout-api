package function

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// ListIndoorSessions returns all indoor sessions, with optional date filtering
func ListIndoorSessions(w http.ResponseWriter, r *http.Request, client *firestore.Client) {
	ctx := context.Background()
	col := GetCollection(client)

	// Optional date range filters
	startDate := r.URL.Query().Get("startDate")
	endDate := r.URL.Query().Get("endDate")

	query := col.OrderBy("date", firestore.Desc)

	if startDate != "" {
		query = query.Where("date", ">=", startDate)
	}
	if endDate != "" {
		query = query.Where("date", "<=", endDate)
	}

	iter := query.Documents(ctx)
	defer iter.Stop()

	var sessions []IndoorSession
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			http.Error(w, "Failed to fetch sessions", http.StatusInternalServerError)
			return
		}

		var session IndoorSession
		if err := doc.DataTo(&session); err != nil {
			continue // Skip malformed documents
		}
		session.ID = doc.Ref.ID
		sessions = append(sessions, session)
	}

	if sessions == nil {
		sessions = []IndoorSession{} // Return empty array, not null
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sessions)
}

// GetIndoorSession returns a single session by ID
func GetIndoorSession(w http.ResponseWriter, r *http.Request, client *firestore.Client, id string) {
	ctx := context.Background()
	col := GetCollection(client)

	doc, err := col.Doc(id).Get(ctx)
	if err != nil {
		http.Error(w, "Session not found", http.StatusNotFound)
		return
	}

	var session IndoorSession
	if err := doc.DataTo(&session); err != nil {
		http.Error(w, "Failed to parse session", http.StatusInternalServerError)
		return
	}
	session.ID = doc.Ref.ID

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(session)
}

// CreateIndoorSession creates a new session
func CreateIndoorSession(w http.ResponseWriter, r *http.Request, client *firestore.Client) {
	ctx := context.Background()
	col := GetCollection(client)

	var input IndoorSessionInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	now := time.Now()
	session := IndoorSession{
		Date:           input.Date,
		Location:       input.Location,
		CustomLocation: input.CustomLocation,
		ClimbingType:   input.ClimbingType,
		SessionType:    input.SessionType,
		FingerLoad:     input.FingerLoad,
		ShoulderLoad:   input.ShoulderLoad,
		ForearmLoad:    input.ForearmLoad,
		Climbs:         input.Climbs,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	docRef, _, err := col.Add(ctx, session)
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	session.ID = docRef.ID

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(session)
}

// UpdateIndoorSession updates an existing session
func UpdateIndoorSession(w http.ResponseWriter, r *http.Request, client *firestore.Client, id string) {
	ctx := context.Background()
	col := GetCollection(client)

	// Check if exists
	docRef := col.Doc(id)
	_, err := docRef.Get(ctx)
	if err != nil {
		http.Error(w, "Session not found", http.StatusNotFound)
		return
	}

	var input IndoorSessionInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updates := []firestore.Update{
		{Path: "date", Value: input.Date},
		{Path: "location", Value: input.Location},
		{Path: "customLocation", Value: input.CustomLocation},
		{Path: "climbingType", Value: input.ClimbingType},
		{Path: "sessionType", Value: input.SessionType},
		{Path: "fingerLoad", Value: input.FingerLoad},
		{Path: "shoulderLoad", Value: input.ShoulderLoad},
		{Path: "forearmLoad", Value: input.ForearmLoad},
		{Path: "climbs", Value: input.Climbs},
		{Path: "updatedAt", Value: time.Now()},
	}

	_, err = docRef.Update(ctx, updates)
	if err != nil {
		http.Error(w, "Failed to update session", http.StatusInternalServerError)
		return
	}

	// Return updated session
	GetIndoorSession(w, r, client, id)
}

// DeleteIndoorSession deletes a session by ID
func DeleteIndoorSession(w http.ResponseWriter, r *http.Request, client *firestore.Client, id string) {
	ctx := context.Background()
	col := GetCollection(client)

	docRef := col.Doc(id)
	_, err := docRef.Get(ctx)
	if err != nil {
		http.Error(w, "Session not found", http.StatusNotFound)
		return
	}

	_, err = docRef.Delete(ctx)
	if err != nil {
		http.Error(w, "Failed to delete session", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ParseSessionID extracts the session ID from path like /indoor_sessions/{id}
func ParseSessionID(path string) string {
	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")
	if len(parts) >= 2 && parts[0] == "indoor_sessions" {
		return parts[1]
	}
	return ""
}
