package function

import (
	"context"
	"sync"

	"cloud.google.com/go/firestore"
)

const (
	// Database and collection names
	DatabaseID     = "climbing-tracker-db"
	CollectionName = "Indoor_Climbs"
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

// GetCollection returns the Indoor_Climbs collection reference
func GetCollection(client *firestore.Client) *firestore.CollectionRef {
	return client.Collection(CollectionName)
}
