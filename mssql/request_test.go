package mssql_test

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.nhat.io/testcontainers-extra"

	"go.nhat.io/testcontainers-registry/mssql"
)

func TestRunMigrations(t *testing.T) {
	t.Parallel()

	dbName := "db"
	dbPassword := "My!StrongPassword"
	migrationSource := "file://./resources/migrations/"

	c, err := mssql.StartGenericContainer(context.Background(),
		dbName, dbPassword,
		mssql.RunMigrations(migrationSource),
		testcontainers.ContainerCallback(func(context.Context, testcontainers.Container, testcontainers.ContainerRequest) error {
			db, err := sql.Open("sqlserver", os.ExpandEnv(mssql.DSN(dbName, dbPassword)))
			if err != nil {
				return err //nolint: wrapcheck
			}

			row := db.QueryRow("SELECT COUNT(1) FROM customer;")
			result := 0

			if err := row.Scan(&result); err != nil {
				return err //nolint: wrapcheck
			}

			assert.Equal(t, 1, result)

			return nil
		}),
	)

	assert.NoError(t, err)
	require.NotNil(t, c)

	_ = testcontainers.StopGenericContainers(context.Background(), c) // nolint: errcheck
}
