package mongo_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/nhatthm/testcontainers-go-extra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	mongodb "go.nhat.io/testcontainers-go-registry/database/mongo"
)

func TestRunMigrations(t *testing.T) {
	t.Parallel()

	dbName := "db"
	migrationSource := "file://./resources/migrations/"

	c, err := mongodb.StartGenericContainer(context.Background(),
		mongodb.RunMigrations(migrationSource, dbName),
		testcontainers.ContainerCallback(func(ctx context.Context, _ testcontainers.Container, _ testcontainers.ContainerRequest) error {
			ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
			defer cancel()

			client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.ExpandEnv(mongodb.DSN(dbName))))
			if err != nil {
				return fmt.Errorf("failed to connect to mongo: %w", err)
			}

			defer client.Disconnect(context.Background()) // nolint: errcheck,contextcheck

			db := client.Database(dbName)
			col := db.Collection("customer")

			result, err := col.CountDocuments(ctx, bson.D{})
			if err != nil {
				return fmt.Errorf("failed to count documents: %w", err)
			}

			assert.Equal(t, int64(1), result)

			return nil
		}),
	)

	assert.NoError(t, err)
	require.NotNil(t, c)

	_ = testcontainers.StopGenericContainers(context.Background(), c) // nolint: errcheck
}
