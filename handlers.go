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
		Date:             input.Date,
		Location:         input.Location,
		CustomLocation:   input.CustomLocation,
		ClimbingType:     input.ClimbingType,
		TrainingTypes:    input.TrainingTypes,
		Difficulty:       input.Difficulty,
		Categories:       input.Categories,
		EnergySystems:    input.EnergySystems,
		TechniqueFocuses: input.TechniqueFocuses,
		FingerLoad:       input.FingerLoad,
		ShoulderLoad:     input.ShoulderLoad,
		ForearmLoad:      input.ForearmLoad,
		Climbs:           input.Climbs,
		CreatedAt:        now,
		UpdatedAt:        now,
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
		{Path: "trainingTypes", Value: input.TrainingTypes},
		{Path: "difficulty", Value: input.Difficulty},
		{Path: "categories", Value: input.Categories},
		{Path: "energySystems", Value: input.EnergySystems},
		{Path: "techniqueFocuses", Value: input.TechniqueFocuses},
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

// ParseOutdoorSessionID extracts the session ID from path like /outdoor_sessions/{id}
func ParseOutdoorSessionID(path string) string {
	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")
	if len(parts) >= 2 && parts[0] == "outdoor_sessions" {
		return parts[1]
	}
	return ""
}

// ListOutdoorSessions returns all outdoor sessions, with optional date filtering
func ListOutdoorSessions(w http.ResponseWriter, r *http.Request, client *firestore.Client) {
	ctx := context.Background()
	col := GetCollectionByName(client, OutdoorCollection)

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

	var sessions []OutdoorSession
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			http.Error(w, "Failed to fetch sessions", http.StatusInternalServerError)
			return
		}

		var session OutdoorSession
		if err := doc.DataTo(&session); err != nil {
			continue
		}
		session.ID = doc.Ref.ID
		sessions = append(sessions, session)
	}

	if sessions == nil {
		sessions = []OutdoorSession{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sessions)
}

// GetOutdoorSession returns a single outdoor session by ID
func GetOutdoorSession(w http.ResponseWriter, r *http.Request, client *firestore.Client, id string) {
	ctx := context.Background()
	col := GetCollectionByName(client, OutdoorCollection)

	doc, err := col.Doc(id).Get(ctx)
	if err != nil {
		http.Error(w, "Session not found", http.StatusNotFound)
		return
	}

	var session OutdoorSession
	if err := doc.DataTo(&session); err != nil {
		http.Error(w, "Failed to parse session", http.StatusInternalServerError)
		return
	}
	session.ID = doc.Ref.ID

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(session)
}

// CreateOutdoorSession creates a new outdoor session
func CreateOutdoorSession(w http.ResponseWriter, r *http.Request, client *firestore.Client) {
	ctx := context.Background()
	col := GetCollectionByName(client, OutdoorCollection)

	var input OutdoorSessionInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	now := time.Now()
	session := OutdoorSession{
		Date:             input.Date,
		Area:             input.Area,
		Crag:             input.Crag,
		Sector:           input.Sector,
		ClimbingType:     input.ClimbingType,
		TrainingTypes:    input.TrainingTypes,
		Difficulty:       input.Difficulty,
		Categories:       input.Categories,
		EnergySystems:    input.EnergySystems,
		TechniqueFocuses: input.TechniqueFocuses,
		FingerLoad:       input.FingerLoad,
		ShoulderLoad:     input.ShoulderLoad,
		ForearmLoad:      input.ForearmLoad,
		Climbs:           input.Climbs,
		CreatedAt:        now,
		UpdatedAt:        now,
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

// UpdateOutdoorSession updates an existing outdoor session
func UpdateOutdoorSession(w http.ResponseWriter, r *http.Request, client *firestore.Client, id string) {
	ctx := context.Background()
	col := GetCollectionByName(client, OutdoorCollection)

	docRef := col.Doc(id)
	_, err := docRef.Get(ctx)
	if err != nil {
		http.Error(w, "Session not found", http.StatusNotFound)
		return
	}

	var input OutdoorSessionInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updates := []firestore.Update{
		{Path: "date", Value: input.Date},
		{Path: "area", Value: input.Area},
		{Path: "crag", Value: input.Crag},
		{Path: "sector", Value: input.Sector},
		{Path: "climbingType", Value: input.ClimbingType},
		{Path: "trainingTypes", Value: input.TrainingTypes},
		{Path: "difficulty", Value: input.Difficulty},
		{Path: "categories", Value: input.Categories},
		{Path: "energySystems", Value: input.EnergySystems},
		{Path: "techniqueFocuses", Value: input.TechniqueFocuses},
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

	GetOutdoorSession(w, r, client, id)
}

// DeleteOutdoorSession deletes an outdoor session by ID
func DeleteOutdoorSession(w http.ResponseWriter, r *http.Request, client *firestore.Client, id string) {
	ctx := context.Background()
	col := GetCollectionByName(client, OutdoorCollection)

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

// ParseFingerboardSessionID extracts the session ID from path like /fingerboard_sessions/{id}
func ParseFingerboardSessionID(path string) string {
	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")
	if len(parts) >= 2 && parts[0] == "fingerboard_sessions" {
		return parts[1]
	}
	return ""
}

// ListFingerboardSessions
func ListFingerboardSessions(w http.ResponseWriter, r *http.Request, client *firestore.Client) {
	ctx := context.Background()
	col := GetCollectionByName(client, FingerboardCollection)

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
	var sessions []FingerboardSession
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			http.Error(w, "Failed to fetch", http.StatusInternalServerError)
			return
		}
		var s FingerboardSession
		if err := doc.DataTo(&s); err == nil {
			s.ID = doc.Ref.ID
			sessions = append(sessions, s)
		}
	}
	if sessions == nil {
		sessions = []FingerboardSession{}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sessions)
}

// GetFingerboardSession
func GetFingerboardSession(w http.ResponseWriter, r *http.Request, client *firestore.Client, id string) {
	ctx := context.Background()
	doc, err := GetCollectionByName(client, FingerboardCollection).Doc(id).Get(ctx)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	var s FingerboardSession
	if err := doc.DataTo(&s); err != nil {
		http.Error(w, "Parse error", http.StatusInternalServerError)
		return
	}
	s.ID = doc.Ref.ID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

// CreateFingerboardSession
func CreateFingerboardSession(w http.ResponseWriter, r *http.Request, client *firestore.Client) {
	ctx := context.Background()
	var input FingerboardSessionInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	now := time.Now()
	s := FingerboardSession{
		Date: input.Date, Location: input.Location, Exercises: input.Exercises, CreatedAt: now, UpdatedAt: now,
	}
	docRef, _, err := GetCollectionByName(client, FingerboardCollection).Add(ctx, s)
	if err != nil {
		http.Error(w, "Failed to create", http.StatusInternalServerError)
		return
	}
	s.ID = docRef.ID
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s)
}

// UpdateFingerboardSession
func UpdateFingerboardSession(w http.ResponseWriter, r *http.Request, client *firestore.Client, id string) {
	ctx := context.Background()
	docRef := GetCollectionByName(client, FingerboardCollection).Doc(id)
	if _, err := docRef.Get(ctx); err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	var input FingerboardSessionInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	updates := []firestore.Update{
		{Path: "date", Value: input.Date},
		{Path: "location", Value: input.Location},
		{Path: "exercises", Value: input.Exercises},
		{Path: "updatedAt", Value: time.Now()},
	}
	if _, err := docRef.Update(ctx, updates); err != nil {
		http.Error(w, "Update failed", http.StatusInternalServerError)
		return
	}
	GetFingerboardSession(w, r, client, id)
}

// DeleteFingerboardSession
func DeleteFingerboardSession(w http.ResponseWriter, r *http.Request, client *firestore.Client, id string) {
	ctx := context.Background()
	docRef := GetCollectionByName(client, FingerboardCollection).Doc(id)
	if _, err := docRef.Delete(ctx); err != nil {
		http.Error(w, "Delete failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// ParseCompetitionSessionID
func ParseCompetitionSessionID(path string) string {
	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")
	if len(parts) >= 2 && parts[0] == "competition_sessions" {
		return parts[1]
	}
	return ""
}

// ListCompetitionSessions
func ListCompetitionSessions(w http.ResponseWriter, r *http.Request, client *firestore.Client) {
	ctx := context.Background()
	col := GetCollectionByName(client, CompetitionCollection)
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
	var sessions []CompetitionSession
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			http.Error(w, "Failed to fetch", http.StatusInternalServerError)
			return
		}
		var s CompetitionSession
		if err := doc.DataTo(&s); err == nil {
			s.ID = doc.Ref.ID
			sessions = append(sessions, s)
		}
	}
	if sessions == nil {
		sessions = []CompetitionSession{}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sessions)
}

// GetCompetitionSession
func GetCompetitionSession(w http.ResponseWriter, r *http.Request, client *firestore.Client, id string) {
	ctx := context.Background()
	doc, err := GetCollectionByName(client, CompetitionCollection).Doc(id).Get(ctx)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	var s CompetitionSession
	if err := doc.DataTo(&s); err != nil {
		http.Error(w, "Parse error", http.StatusInternalServerError)
		return
	}
	s.ID = doc.Ref.ID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

// CreateCompetitionSession
func CreateCompetitionSession(w http.ResponseWriter, r *http.Request, client *firestore.Client) {
	ctx := context.Background()
	var input CompetitionSessionInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	now := time.Now()
	s := CompetitionSession{
		Date: input.Date, Venue: input.Venue, CustomVenue: input.CustomVenue, Type: input.Type,
		FingerLoad: input.FingerLoad, ShoulderLoad: input.ShoulderLoad, ForearmLoad: input.ForearmLoad,
		Rounds: input.Rounds, IsSimulation: input.IsSimulation, CreatedAt: now, UpdatedAt: now,
	}
	docRef, _, err := GetCollectionByName(client, CompetitionCollection).Add(ctx, s)
	if err != nil {
		http.Error(w, "Failed to create", http.StatusInternalServerError)
		return
	}
	s.ID = docRef.ID
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s)
}

// UpdateCompetitionSession
func UpdateCompetitionSession(w http.ResponseWriter, r *http.Request, client *firestore.Client, id string) {
	ctx := context.Background()
	docRef := GetCollectionByName(client, CompetitionCollection).Doc(id)
	if _, err := docRef.Get(ctx); err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	var input CompetitionSessionInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	updates := []firestore.Update{
		{Path: "date", Value: input.Date},
		{Path: "venue", Value: input.Venue},
		{Path: "customVenue", Value: input.CustomVenue},
		{Path: "type", Value: input.Type},
		{Path: "fingerLoad", Value: input.FingerLoad},
		{Path: "shoulderLoad", Value: input.ShoulderLoad},
		{Path: "forearmLoad", Value: input.ForearmLoad},
		{Path: "rounds", Value: input.Rounds},
		{Path: "isSimulation", Value: input.IsSimulation},
		{Path: "updatedAt", Value: time.Now()},
	}
	if _, err := docRef.Update(ctx, updates); err != nil {
		http.Error(w, "Update failed", http.StatusInternalServerError)
		return
	}
	GetCompetitionSession(w, r, client, id)
}

// DeleteCompetitionSession
func DeleteCompetitionSession(w http.ResponseWriter, r *http.Request, client *firestore.Client, id string) {
	ctx := context.Background()
	docRef := GetCollectionByName(client, CompetitionCollection).Doc(id)
	if _, err := docRef.Delete(ctx); err != nil {
		http.Error(w, "Delete failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// ParseGymSessionID wrapper
func ParseGymSessionID(path string) string {
	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")
	if len(parts) >= 2 && parts[0] == "gym_sessions" {
		return parts[1]
	}
	return ""
}

// ListGymSessions
func ListGymSessions(w http.ResponseWriter, r *http.Request, client *firestore.Client) {
	ctx := context.Background()
	col := GetCollectionByName(client, GymCollection)
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
	var sessions []GymSession
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			http.Error(w, "Failed to fetch", http.StatusInternalServerError)
			return
		}
		var s GymSession
		if err := doc.DataTo(&s); err == nil {
			s.ID = doc.Ref.ID
			sessions = append(sessions, s)
		}
	}
	if sessions == nil {
		sessions = []GymSession{}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sessions)
}

// GetGymSession
func GetGymSession(w http.ResponseWriter, r *http.Request, client *firestore.Client, id string) {
	ctx := context.Background()
	doc, err := GetCollectionByName(client, GymCollection).Doc(id).Get(ctx)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	var s GymSession
	if err := doc.DataTo(&s); err != nil {
		http.Error(w, "Parse error", http.StatusInternalServerError)
		return
	}
	s.ID = doc.Ref.ID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

// CreateGymSession
func CreateGymSession(w http.ResponseWriter, r *http.Request, client *firestore.Client) {
	ctx := context.Background()
	var input GymSessionInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	now := time.Now()
	s := GymSession{
		Date: input.Date, Name: input.Name, Bodyweight: input.Bodyweight, TrainingBlock: input.TrainingBlock, Exercises: input.Exercises, CreatedAt: now, UpdatedAt: now,
	}
	docRef, _, err := GetCollectionByName(client, GymCollection).Add(ctx, s)
	if err != nil {
		http.Error(w, "Failed to create", http.StatusInternalServerError)
		return
	}
	s.ID = docRef.ID
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s)
}

// UpdateGymSession
func UpdateGymSession(w http.ResponseWriter, r *http.Request, client *firestore.Client, id string) {
	ctx := context.Background()
	docRef := GetCollectionByName(client, GymCollection).Doc(id)
	if _, err := docRef.Get(ctx); err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	var input GymSessionInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	updates := []firestore.Update{
		{Path: "date", Value: input.Date},
		{Path: "name", Value: input.Name},
		{Path: "bodyweight", Value: input.Bodyweight},
		{Path: "trainingBlock", Value: input.TrainingBlock},
		{Path: "exercises", Value: input.Exercises},
		{Path: "updatedAt", Value: time.Now()},
	}
	if _, err := docRef.Update(ctx, updates); err != nil {
		http.Error(w, "Update failed", http.StatusInternalServerError)
		return
	}
	GetGymSession(w, r, client, id)
}

// DeleteGymSession
func DeleteGymSession(w http.ResponseWriter, r *http.Request, client *firestore.Client, id string) {
	ctx := context.Background()
	docRef := GetCollectionByName(client, GymCollection).Doc(id)
	if _, err := docRef.Delete(ctx); err != nil {
		http.Error(w, "Delete failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
