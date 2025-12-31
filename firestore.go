package function

import (
	"context"
	"sync"

	"cloud.google.com/go/firestore"
)

const (
	// Database name
	DatabaseID = "climbing-tracker-db"

	// Collection names
	IndoorCollection      = "Indoor_Climbs"
	OutdoorCollection     = "Outdoor_Climbs"
	FingerboardCollection = "Fingerboarding"
	CompetitionCollection = "Competitions"
	GymCollection         = "Gym_Sessions"
)

var (
	firestoreClient *firestore.Client
	clientOnce      sync.Once
	clientErr       error
)

// GetFirestoreClient returns a singleton Firestore client
func GetFirestoreClient(ctx context.Context, projectID string) (*firestore.Client, error) {
	clientOnce.Do(func() {
		firestoreClient, clientErr = firestore.NewClientWithDatabase(ctx, projectID, DatabaseID)
	})
	return firestoreClient, clientErr
}

// GetCollection returns the Indoor_Climbs collection reference (legacy support)
func GetCollection(client *firestore.Client) *firestore.CollectionRef {
	return client.Collection(IndoorCollection)
}

// GetCollectionByName returns a collection reference by name
func GetCollectionByName(client *firestore.Client, name string) *firestore.CollectionRef {
	return client.Collection(name)
}
