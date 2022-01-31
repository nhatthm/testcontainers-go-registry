package testcontainerspostgres

import (
	"fmt"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres" // Database driver
	"github.com/nhatthm/testcontainers-go-extra"
	"github.com/nhatthm/testcontainers-go-extra/wait"

	testcontainersql "github.com/nhatthm/testcontainers-go-registry/sql"
)

// Request creates a new request for starting a postgres server.
func Request(dbName, dbUser, dbPassword string, opts ...testcontainers.GenericContainerOption) testcontainers.StartGenericContainerRequest {
	finalOpts := make([]testcontainers.GenericContainerOption, 1, len(opts)+1)
	finalOpts[0] = testcontainers.PopulateHostPortEnv
	finalOpts = append(finalOpts, opts...)

	return testcontainers.StartGenericContainerRequest{
		Request: testcontainers.ContainerRequest{
			Name:         "postgres",
			Image:        "postgres:12-alpine",
			ExposedPorts: []string{":5432"},
			Env: map[string]string{
				"LC_ALL":            "C.UTF-8",
				"POSTGRES_DB":       dbName,
				"POSTGRES_USER":     dbUser,
				"POSTGRES_PASSWORD": dbPassword,
			},
			WaitingFor: wait.ForHealthCheckCmd("pg_isready").
				WithRetries(3).
				WithStartPeriod(15 * time.Second).
				WithTestTimeout(5 * time.Second).
				WithTestInterval(10 * time.Second),
		},
		Options: finalOpts,
	}
}

// RunMigrations is option to run migrations.
func RunMigrations(migrationSource string) testcontainers.ContainerCallback {
	return testcontainersql.RunMigrations(migrationSource, DSN(`$POSTGRES_DB`, `$POSTGRES_USER`, `$POSTGRES_PASSWORD`))
}

// DSN returns the database dsn for connecting to server.
func DSN(dbName, dbUser, dbPassword string) string {
	return fmt.Sprintf("postgres://%s:%s@$POSTGRES_5432_HOST:$POSTGRES_5432_PORT/%s?sslmode=disable&client_encoding=UTF8", dbUser, dbPassword, dbName)
}
