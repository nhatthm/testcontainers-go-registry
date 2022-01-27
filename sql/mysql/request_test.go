package testcontainersmysql_test

import (
	"context"
	"database/sql"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql" // Database driver
	"github.com/nhatthm/testcontainers-go-extra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	testcontainermysql "github.com/nhatthm/testcontainers-go-registry/sql/mysql"
)

func TestRunMigrations(t *testing.T) {
	t.Parallel()

	dbName := "test"
	dbUser := "test"
	dbPassword := "test"
	migrationSource := "file://./resources/migrations/"

	c, err := testcontainermysql.StartGenericContainer(context.Background(),
		dbName, dbUser, dbPassword,
		testcontainermysql.RunMigrations(migrationSource),
		testcontainers.ContainerCallback(func(context.Context, testcontainers.Container, testcontainers.ContainerRequest) error {
			db, err := sql.Open("mysql", os.ExpandEnv(testcontainermysql.DSN(dbName, dbUser, dbPassword)))
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
