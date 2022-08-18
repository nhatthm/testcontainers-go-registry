package mongo

import (
	"fmt"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/mongodb" // Database driver
	"go.nhat.io/testcontainers-go-extra"
	"go.nhat.io/testcontainers-go-extra/wait"

	db "go.nhat.io/testcontainers-go-registry/database"
)

// Request creates a new request for starting a mongo server.
func Request(opts ...testcontainers.GenericContainerOption) testcontainers.StartGenericContainerRequest {
	finalOpts := make([]testcontainers.GenericContainerOption, 1, len(opts)+1)
	finalOpts[0] = testcontainers.PopulateHostPortEnv
	finalOpts = append(finalOpts, opts...)

	return testcontainers.StartGenericContainerRequest{
		Request: testcontainers.ContainerRequest{
			Name:         "mongo",
			Image:        "mongo:4.4",
			ExposedPorts: []string{":27017"},
			WaitingFor: wait.ForHealthCheckCmd("mongo", "localhost:27017/test", "--quiet", "--eval", "'quit(db.runCommand({ ping: 1 }).ok ? 0 : 2)'").
				WithRetries(3).
				WithStartPeriod(15 * time.Second).
				WithTestTimeout(5 * time.Second).
				WithTestInterval(10 * time.Second),
		},
		Options: finalOpts,
	}
}

// RunMigrations is option to run migrations.
func RunMigrations(migrationSource, dbName string) testcontainers.ContainerCallback {
	return db.RunMigrations(migrationSource, DSN(dbName))
}

// DSN returns the database dsn for connecting to server.
func DSN(dbName string) string {
	return fmt.Sprintf("mongodb://$MONGO_27017_HOST:$MONGO_27017_PORT/%s", dbName)
}
