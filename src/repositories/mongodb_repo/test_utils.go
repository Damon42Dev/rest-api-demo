package mongodb_repo

import (
	"context"
	"os/exec"
	"testing"
	"time"

	"example/rest-api-demo/src/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TestDBClient holds the MongoDB client and repositories for testing
type TestDBClient struct {
	Client *mongo.Client
	Config *utils.Configuration
}

// waitForMongoDB attempts to connect to MongoDB with retries
func waitForMongoDB(ctx context.Context, uri string) (*mongo.Client, error) {
	opt := options.Client().
		ApplyURI(uri).
		SetServerSelectionTimeout(5 * time.Second).
		SetConnectTimeout(5 * time.Second)

	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		return nil, err
	}

	// Try to ping the database
	if err = client.Ping(ctx, nil); err != nil {
		client.Disconnect(ctx)
		return nil, err
	}

	return client, nil
}

// SetupTestDB sets up a test database connection and configuration
func SetupTestDB(t *testing.T) *TestDBClient {
	t.Helper()
	ctx := context.Background()

	// Run the setup script
	cmd := exec.Command("bash", "../../../scripts/setup_test_db.sh")
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to run setup script: %v", err)
	}

	// Wait a short time for MongoDB to be ready
	time.Sleep(2 * time.Second)

	// Connect to MongoDB
	client, err := waitForMongoDB(ctx, "mongodb://localhost:27017")
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Create a test configuration
	config := &utils.Configuration{
		Database: utils.DatabaseSetting{
			DbName:      "test_db",
			Collections: []string{"comments", "users", "movies"},
		},
	}

	return &TestDBClient{
		Client: client,
		Config: config,
	}
}

// TeardownTestDB cleans up the test database and connection
func TeardownTestDB(t *testing.T, client *mongo.Client) {
	t.Helper()
	ctx := context.Background()

	if client != nil {
		if err := client.Disconnect(ctx); err != nil {
			t.Errorf("Failed to disconnect from MongoDB: %v", err)
		}
	}

	cmd := exec.Command("bash", "../../../scripts/teardown_test_db.sh")
	if err := cmd.Run(); err != nil {
		t.Errorf("Failed to run teardown script: %v", err)
	}
}
