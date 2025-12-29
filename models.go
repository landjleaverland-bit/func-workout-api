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
	TrainingType   string       `json:"trainingType" firestore:"trainingType"`
	Difficulty     string       `json:"difficulty,omitempty" firestore:"difficulty,omitempty"`
	Category       string       `json:"category,omitempty" firestore:"category,omitempty"`
	EnergySystem   string       `json:"energySystem,omitempty" firestore:"energySystem,omitempty"`
	TechniqueFocus string       `json:"techniqueFocus,omitempty" firestore:"techniqueFocus,omitempty"`
	WallAngle      string       `json:"wallAngle,omitempty" firestore:"wallAngle,omitempty"`
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
	TrainingType   string       `json:"trainingType"`
	Difficulty     string       `json:"difficulty,omitempty"`
	Category       string       `json:"category,omitempty"`
	EnergySystem   string       `json:"energySystem,omitempty"`
	TechniqueFocus string       `json:"techniqueFocus,omitempty"`
	WallAngle      string       `json:"wallAngle,omitempty"`
	FingerLoad     int          `json:"fingerLoad"`
	ShoulderLoad   int          `json:"shoulderLoad"`
	ForearmLoad    int          `json:"forearmLoad"`
	Climbs         []ClimbEntry `json:"climbs"`
}

// OutdoorSession represents an outdoor climbing session
type OutdoorSession struct {
	ID             string       `json:"id" firestore:"-"`
	Date           string       `json:"date" firestore:"date"`
	Area           string       `json:"area" firestore:"area"`
	Crag           string       `json:"crag" firestore:"crag"`
	Sector         string       `json:"sector,omitempty" firestore:"sector,omitempty"`
	ClimbingType   string       `json:"climbingType" firestore:"climbingType"`
	TrainingType   string       `json:"trainingType" firestore:"trainingType"`
	Difficulty     string       `json:"difficulty,omitempty" firestore:"difficulty,omitempty"`
	Category       string       `json:"category,omitempty" firestore:"category,omitempty"`
	EnergySystem   string       `json:"energySystem,omitempty" firestore:"energySystem,omitempty"`
	TechniqueFocus string       `json:"techniqueFocus,omitempty" firestore:"techniqueFocus,omitempty"`
	FingerLoad     int          `json:"fingerLoad" firestore:"fingerLoad"`
	ShoulderLoad   int          `json:"shoulderLoad" firestore:"shoulderLoad"`
	ForearmLoad    int          `json:"forearmLoad" firestore:"forearmLoad"`
	Climbs         []ClimbEntry `json:"climbs" firestore:"climbs"`
	CreatedAt      time.Time    `json:"createdAt" firestore:"createdAt"`
	UpdatedAt      time.Time    `json:"updatedAt" firestore:"updatedAt"`
}

// OutdoorSessionInput is used for create/update outdoor session requests
type OutdoorSessionInput struct {
	Date           string       `json:"date"`
	Area           string       `json:"area"`
	Crag           string       `json:"crag"`
	Sector         string       `json:"sector,omitempty"`
	ClimbingType   string       `json:"climbingType"`
	TrainingType   string       `json:"trainingType"`
	Difficulty     string       `json:"difficulty,omitempty"`
	Category       string       `json:"category,omitempty"`
	EnergySystem   string       `json:"energySystem,omitempty"`
	TechniqueFocus string       `json:"techniqueFocus,omitempty"`
	FingerLoad     int          `json:"fingerLoad"`
	ShoulderLoad   int          `json:"shoulderLoad"`
	ForearmLoad    int          `json:"forearmLoad"`
	Climbs         []ClimbEntry `json:"climbs"`
}
