package mssql

import (
	"context"
	"fmt"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/sqlserver" // Database driver
	"github.com/testcontainers/testcontainers-go/wait"
	"go.nhat.io/testcontainers-extra"
	extrawait "go.nhat.io/testcontainers-extra/wait"

	db "go.nhat.io/testcontainers-registry/database"
)

// Request creates a new request for starting a mssql server.
func Request(dbName, dbPassword string, opts ...testcontainers.GenericContainerOption) testcontainers.StartGenericContainerRequest {
	finalOpts := make([]testcontainers.GenericContainerOption, 2, len(opts)+2)

	finalOpts[0] = testcontainers.PopulateHostPortEnv
	finalOpts[1] = testcontainers.WithCallback(func(ctx context.Context, c testcontainers.Container, _ testcontainers.ContainerRequest) error {
		code, _, err := c.Exec(ctx, []string{
			"/opt/mssql-tools/bin/sqlcmd", "-S", "localhost", "-U", "sa", "-P", dbPassword, "-Q", `USE [master]; CREATE DATABASE ` + dbName + `;`,
		})
		if err != nil {
			return fmt.Errorf("could not create database: %w", err)
		}

		if code > 0 {
			return fmt.Errorf("could not create database (code %d)", code) // nolint: goerr113
		}

		return nil
	})

	finalOpts = append(finalOpts, opts...)

	return testcontainers.StartGenericContainerRequest{
		Request: testcontainers.ContainerRequest{
			Name:         "mssql",
			Image:        "mcr.microsoft.com/mssql/server:2019-latest",
			ExposedPorts: []string{":1433"},
			Env: map[string]string{
				"LC_ALL":      "C.UTF-8",
				"ACCEPT_EULA": "Y",
				"SA_PASSWORD": dbPassword,
				"MSSQL_PID":   "Developer",
				"MSSQL_DB":    dbName,
			},
			WaitingFor: wait.ForAll(
				wait.ForLog("Recovery is complete").
					WithStartupTimeout(3*time.Minute),
				extrawait.ForHealthCheckCmd("/opt/mssql-tools/bin/sqlcmd", "-S", "localhost", "-U", "sa", "-P", dbPassword, "-Q", `"SELECT 1"`).
					WithRetries(3).
					WithStartPeriod(5*time.Minute).
					WithTestTimeout(5*time.Second).
					WithTestInterval(10*time.Second),
			).WithStartupTimeout(10 * time.Minute),
		},
		Options: finalOpts,
	}
}

// RunMigrations is option to run migrations.
func RunMigrations(migrationSource string) testcontainers.ContainerCallback {
	return db.RunMigrations(migrationSource, DSN(`$MSSQL_DB`, `$SA_PASSWORD`))
}

// DSN returns the database dsn for connecting to server.
func DSN(dbName, dbPassword string) string {
	return fmt.Sprintf("sqlserver://sa:%s@$MSSQL_1433_HOST:$MSSQL_1433_PORT?database=%s", dbPassword, dbName)
}
