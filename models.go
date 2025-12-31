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
	ID               string       `json:"id" firestore:"-"`
	Date             string       `json:"date" firestore:"date"`
	Location         string       `json:"location" firestore:"location"`
	CustomLocation   string       `json:"customLocation,omitempty" firestore:"customLocation,omitempty"`
	ClimbingType     string       `json:"climbingType" firestore:"climbingType"`
	TrainingTypes    []string     `json:"trainingTypes" firestore:"trainingTypes"`
	Difficulty       string       `json:"difficulty,omitempty" firestore:"difficulty,omitempty"`
	Categories       []string     `json:"categories,omitempty" firestore:"categories,omitempty"`
	EnergySystems    []string     `json:"energySystems,omitempty" firestore:"energySystems,omitempty"`
	TechniqueFocuses []string     `json:"techniqueFocuses,omitempty" firestore:"techniqueFocuses,omitempty"`
	WallAngles       []string     `json:"wallAngles,omitempty" firestore:"wallAngles,omitempty"`
	FingerLoad       int          `json:"fingerLoad" firestore:"fingerLoad"`
	ShoulderLoad     int          `json:"shoulderLoad" firestore:"shoulderLoad"`
	ForearmLoad      int          `json:"forearmLoad" firestore:"forearmLoad"`
	OpenGrip         int          `json:"openGrip" firestore:"openGrip"`
	CrimpGrip        int          `json:"crimpGrip" firestore:"crimpGrip"`
	PinchGrip        int          `json:"pinchGrip" firestore:"pinchGrip"`
	SloperGrip       int          `json:"sloperGrip" firestore:"sloperGrip"`
	JugGrip          int          `json:"jugGrip" firestore:"jugGrip"`
	Climbs           []ClimbEntry `json:"climbs" firestore:"climbs"`
	CreatedAt        time.Time    `json:"createdAt" firestore:"createdAt"`
	UpdatedAt        time.Time    `json:"updatedAt" firestore:"updatedAt"`
}

// IndoorSessionInput is used for create/update requests (no ID/timestamps)
type IndoorSessionInput struct {
	Date             string       `json:"date"`
	Location         string       `json:"location"`
	CustomLocation   string       `json:"customLocation,omitempty"`
	ClimbingType     string       `json:"climbingType"`
	TrainingTypes    []string     `json:"trainingTypes"`
	Difficulty       string       `json:"difficulty,omitempty"`
	Categories       []string     `json:"categories,omitempty"`
	EnergySystems    []string     `json:"energySystems,omitempty"`
	TechniqueFocuses []string     `json:"techniqueFocuses,omitempty"`
	WallAngles       []string     `json:"wallAngles,omitempty"`
	FingerLoad       int          `json:"fingerLoad"`
	ShoulderLoad     int          `json:"shoulderLoad"`
	ForearmLoad      int          `json:"forearmLoad"`
	OpenGrip         int          `json:"openGrip"`
	CrimpGrip        int          `json:"crimpGrip"`
	PinchGrip        int          `json:"pinchGrip"`
	SloperGrip       int          `json:"sloperGrip"`
	JugGrip          int          `json:"jugGrip"`
	Climbs           []ClimbEntry `json:"climbs"`
}

// OutdoorSession represents an outdoor climbing session
type OutdoorSession struct {
	ID               string       `json:"id" firestore:"-"`
	Date             string       `json:"date" firestore:"date"`
	Area             string       `json:"area" firestore:"area"`
	Crag             string       `json:"crag" firestore:"crag"`
	Sector           string       `json:"sector,omitempty" firestore:"sector,omitempty"`
	ClimbingType     string       `json:"climbingType" firestore:"climbingType"`
	TrainingTypes    []string     `json:"trainingTypes" firestore:"trainingTypes"`
	Difficulty       string       `json:"difficulty,omitempty" firestore:"difficulty,omitempty"`
	Categories       []string     `json:"categories,omitempty" firestore:"categories,omitempty"`
	EnergySystems    []string     `json:"energySystems,omitempty" firestore:"energySystems,omitempty"`
	TechniqueFocuses []string     `json:"techniqueFocuses,omitempty" firestore:"techniqueFocuses,omitempty"`
	FingerLoad       int          `json:"fingerLoad" firestore:"fingerLoad"`
	ShoulderLoad     int          `json:"shoulderLoad" firestore:"shoulderLoad"`
	ForearmLoad      int          `json:"forearmLoad" firestore:"forearmLoad"`
	OpenGrip         int          `json:"openGrip" firestore:"openGrip"`
	CrimpGrip        int          `json:"crimpGrip" firestore:"crimpGrip"`
	PinchGrip        int          `json:"pinchGrip" firestore:"pinchGrip"`
	SloperGrip       int          `json:"sloperGrip" firestore:"sloperGrip"`
	JugGrip          int          `json:"jugGrip" firestore:"jugGrip"`
	Climbs           []ClimbEntry `json:"climbs" firestore:"climbs"`
	CreatedAt        time.Time    `json:"createdAt" firestore:"createdAt"`
	UpdatedAt        time.Time    `json:"updatedAt" firestore:"updatedAt"`
}

// OutdoorSessionInput is used for create/update outdoor session requests
type OutdoorSessionInput struct {
	Date             string       `json:"date"`
	Area             string       `json:"area"`
	Crag             string       `json:"crag"`
	Sector           string       `json:"sector,omitempty"`
	ClimbingType     string       `json:"climbingType"`
	TrainingTypes    []string     `json:"trainingTypes"`
	Difficulty       string       `json:"difficulty,omitempty"`
	Categories       []string     `json:"categories,omitempty"`
	EnergySystems    []string     `json:"energySystems,omitempty"`
	TechniqueFocuses []string     `json:"techniqueFocuses,omitempty"`
	FingerLoad       int          `json:"fingerLoad"`
	ShoulderLoad     int          `json:"shoulderLoad"`
	ForearmLoad      int          `json:"forearmLoad"`
	OpenGrip         int          `json:"openGrip"`
	CrimpGrip        int          `json:"crimpGrip"`
	PinchGrip        int          `json:"pinchGrip"`
	SloperGrip       int          `json:"sloperGrip"`
	JugGrip          int          `json:"jugGrip"`
	Climbs           []ClimbEntry `json:"climbs"`
}

// Fingerboard Exercise Details
type ExerciseSet struct {
	Weight float64 `json:"weight" firestore:"weight"`
	Reps   int     `json:"reps" firestore:"reps"`
}

type FingerboardExercise struct {
	ID       string        `json:"id" firestore:"id"`
	Name     string        `json:"name" firestore:"name"`
	GripType string        `json:"gripType" firestore:"gripType"`
	Sets     int           `json:"sets" firestore:"sets"`
	Details  []ExerciseSet `json:"details" firestore:"details"`
	Notes    string        `json:"notes" firestore:"notes"`
}

// Fingerboard Session
type FingerboardSession struct {
	ID        string                `json:"id" firestore:"-"`
	Date      string                `json:"date" firestore:"date"`
	Location  string                `json:"location" firestore:"location"` // Usually "N/A" or "Home"
	Exercises []FingerboardExercise `json:"exercises" firestore:"exercises"`
	CreatedAt time.Time             `json:"createdAt" firestore:"createdAt"`
	UpdatedAt time.Time             `json:"updatedAt" firestore:"updatedAt"`
}

type FingerboardSessionInput struct {
	Date      string                `json:"date"`
	Location  string                `json:"location"`
	Exercises []FingerboardExercise `json:"exercises"`
}

// Competition Data
type CompetitionRound struct {
	Name     string                   `json:"name" firestore:"name"` // Qualifiers, Finals, etc.
	Position *int                     `json:"position,omitempty" firestore:"position,omitempty"`
	Climbs   []CompetitionClimbResult `json:"climbs,omitempty" firestore:"climbs,omitempty"`
}

type CompetitionClimbResult struct {
	Name         string `json:"name" firestore:"name"`     // Problem #
	Status       string `json:"status" firestore:"status"` // Flash, Top, Zone, Attempt
	AttemptCount int    `json:"attemptCount" firestore:"attemptCount"`
	Notes        string `json:"notes" firestore:"notes"`
}

// Competition Session
type CompetitionSession struct {
	ID           string             `json:"id" firestore:"-"`
	Date         string             `json:"date" firestore:"date"`
	Venue        string             `json:"venue" firestore:"venue"`
	CustomVenue  string             `json:"customVenue,omitempty" firestore:"customVenue,omitempty"`
	Type         string             `json:"type" firestore:"type"` // Bouldering, Lead, Speed
	FingerLoad   int                `json:"fingerLoad,omitempty" firestore:"fingerLoad,omitempty"`
	ShoulderLoad int                `json:"shoulderLoad,omitempty" firestore:"shoulderLoad,omitempty"`
	ForearmLoad  int                `json:"forearmLoad,omitempty" firestore:"forearmLoad,omitempty"`
	Rounds       []CompetitionRound `json:"rounds" firestore:"rounds"`
	CreatedAt    time.Time          `json:"createdAt" firestore:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt" firestore:"updatedAt"`
}

type CompetitionSessionInput struct {
	Date         string             `json:"date"`
	Venue        string             `json:"venue"`
	CustomVenue  string             `json:"customVenue,omitempty"`
	Type         string             `json:"type"`
	FingerLoad   int                `json:"fingerLoad,omitempty"`
	ShoulderLoad int                `json:"shoulderLoad,omitempty"`
	ForearmLoad  int                `json:"forearmLoad,omitempty"`
	Rounds       []CompetitionRound `json:"rounds"`
}

// Gym Session Data
type GymSet struct {
	Weight    float64 `json:"weight" firestore:"weight"`
	Reps      int     `json:"reps" firestore:"reps"`
	IsWarmup  bool    `json:"isWarmup" firestore:"isWarmup"`
	IsFailure bool    `json:"isFailure" firestore:"isFailure"`
	IsDropSet bool    `json:"isDropSet" firestore:"isDropSet"`
	Completed bool    `json:"completed" firestore:"completed"`
}

type GymExercise struct {
	ID       string   `json:"id" firestore:"id"`
	Name     string   `json:"name" firestore:"name"`
	Sets     []GymSet `json:"sets" firestore:"sets"`
	Notes    string   `json:"notes,omitempty" firestore:"notes,omitempty"`
	LinkedTo string   `json:"linkedTo,omitempty" firestore:"linkedTo,omitempty"`
}

type GymSession struct {
	ID         string        `json:"id" firestore:"-"`
	Date       string        `json:"date" firestore:"date"`
	Name       string        `json:"name" firestore:"name"` // e.g. "Leg Day"
	Bodyweight float64       `json:"bodyweight,omitempty" firestore:"bodyweight,omitempty"`
	Exercises  []GymExercise `json:"exercises" firestore:"exercises"`
	CreatedAt  time.Time     `json:"createdAt" firestore:"createdAt"`
	UpdatedAt  time.Time     `json:"updatedAt" firestore:"updatedAt"`
}

type GymSessionInput struct {
	Date       string        `json:"date"`
	Name       string        `json:"name"`
	Bodyweight float64       `json:"bodyweight,omitempty"`
	Exercises  []GymExercise `json:"exercises"`
}
