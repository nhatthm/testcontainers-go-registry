package postgres_test

import (
	"context"
	"database/sql"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib" // Database driver
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.nhat.io/testcontainers-go-extra"

	pg "go.nhat.io/testcontainers-go-registry/postgres"
)

func TestRunMigrations(t *testing.T) {
	t.Parallel()

	dbName := "db"
	dbUser := "user"
	dbPassword := "password"
	migrationSource := "file://./resources/migrations/"

	c, err := pg.StartGenericContainer(context.Background(),
		dbName, dbUser, dbPassword,
		pg.RunMigrations(migrationSource),
		testcontainers.ContainerCallback(func(context.Context, testcontainers.Container, testcontainers.ContainerRequest) error {
			db, err := sql.Open("pgx", os.ExpandEnv(pg.DSN(dbName, dbUser, dbPassword)))
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

	assert.NoError(t, err)
	require.NotNil(t, c)

	_ = testcontainers.StopGenericContainers(context.Background(), c) // nolint: errcheck
}
