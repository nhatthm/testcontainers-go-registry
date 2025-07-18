package mysql_test

import (
	"context"
	"database/sql"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql" // Database driver
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.nhat.io/testcontainers-extra"

	"go.nhat.io/testcontainers-registry/mysql"
)

func TestRunMigrations(t *testing.T) {
	t.Parallel()

	dbName := "db"
	dbUser := "user"
	dbPassword := "password"
	migrationSource := "file://./resources/migrations/"

	c, err := mysql.StartGenericContainer(context.Background(),
		dbName, dbUser, dbPassword,
		mysql.RunMigrations(migrationSource),
		testcontainers.ContainerCallback(func(context.Context, testcontainers.Container, testcontainers.ContainerRequest) error {
			db, err := sql.Open("mysql", os.ExpandEnv(mysql.DSN(dbName, dbUser, dbPassword)))
			if err != nil {
				return err //nolint: wrapcheck
			}

			row := db.QueryRow("SELECT COUNT(1) FROM customer LIMIT 1")
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
