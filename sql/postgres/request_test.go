package testcontainerspostgres_test

import (
	"context"
	"database/sql"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib" // Database driver
	"github.com/nhatthm/testcontainers-go-extra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	testcontainerpostgres "github.com/nhatthm/testcontainers-go-registry/sql/postgres"
)

func TestRunMigrations(t *testing.T) {
	t.Parallel()

	dbName := "test"
	dbUser := "test"
	dbPassword := "test"
	migrationSource := "file://./resources/migrations/"

	c, err := testcontainerpostgres.StartGenericContainer(context.Background(),
		dbName, dbUser, dbPassword,
		testcontainerpostgres.RunMigrations(migrationSource),
		testcontainers.ContainerCallback(func(context.Context, testcontainers.Container, testcontainers.ContainerRequest) error {
			db, err := sql.Open("pgx", os.ExpandEnv(testcontainerpostgres.DSN(dbName, dbUser, dbPassword)))
			if err != nil {
				return err
			}

			row := db.QueryRow("SELECT COUNT(1) FROM customer LIMIT 1")
			result := 0

			if err := row.Scan(&result); err != nil {
				return err
			}

			assert.Equal(t, 1, result)

			return nil
		}),
	)

	require.NotNil(t, c)
	require.NoError(t, err)

	_ = testcontainers.StopGenericContainers(context.Background(), c) // nolint: errcheck
}
