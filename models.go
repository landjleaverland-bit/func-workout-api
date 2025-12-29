package function

import "time"

// ClimbEntry represents a single climb in a session
type ClimbEntry struct {
	IsSport     bool   `json:"isSport" firestore:"isSport"`
	Name        string `json:"name" firestore:"name"`
	Grade       string `json:"grade" firestore:"grade"`
	AttemptType string `json:"attemptType" firestore:"attemptType"`
	AttemptsNum int    `json:"attemptsNum" firestore:"attemptsNum"`
	Notes       string `json:"notes" firestore:"notes"`
}

// IndoorSession represents an indoor climbing session
type IndoorSession struct {
	ID             string       `json:"id" firestore:"-"`
	Date           string       `json:"date" firestore:"date"`
	Location       string       `json:"location" firestore:"location"`
	CustomLocation string       `json:"customLocation,omitempty" firestore:"customLocation,omitempty"`
	ClimbingType   string       `json:"climbingType" firestore:"climbingType"`
	SessionType    string       `json:"sessionType" firestore:"sessionType"`
	FingerLoad     int          `json:"fingerLoad" firestore:"fingerLoad"`
	ShoulderLoad   int          `json:"shoulderLoad" firestore:"shoulderLoad"`
	ForearmLoad    int          `json:"forearmLoad" firestore:"forearmLoad"`
	Climbs         []ClimbEntry `json:"climbs" firestore:"climbs"`
	CreatedAt      time.Time    `json:"createdAt" firestore:"createdAt"`
	UpdatedAt      time.Time    `json:"updatedAt" firestore:"updatedAt"`
}

// IndoorSessionInput is used for create/update requests (no ID/timestamps)
type IndoorSessionInput struct {
	Date           string       `json:"date"`
	Location       string       `json:"location"`
	CustomLocation string       `json:"customLocation,omitempty"`
	ClimbingType   string       `json:"climbingType"`
	SessionType    string       `json:"sessionType"`
	FingerLoad     int          `json:"fingerLoad"`
	ShoulderLoad   int          `json:"shoulderLoad"`
	ForearmLoad    int          `json:"forearmLoad"`
	Climbs         []ClimbEntry `json:"climbs"`
}
