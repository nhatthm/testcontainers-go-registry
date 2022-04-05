package mysql

import (
	"fmt"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/mysql" // Database driver
	"github.com/nhatthm/testcontainers-go-extra"
	extrawait "github.com/nhatthm/testcontainers-go-extra/wait"
	"github.com/testcontainers/testcontainers-go/wait"

	db "github.com/nhatthm/testcontainers-go-registry/database"
)

// Request creates a new request for starting a mysql server.
func Request(dbName, dbUser, dbPassword string, opts ...testcontainers.GenericContainerOption) testcontainers.StartGenericContainerRequest {
	finalOpts := make([]testcontainers.GenericContainerOption, 1, len(opts)+1)
	finalOpts[0] = testcontainers.PopulateHostPortEnv
	finalOpts = append(finalOpts, opts...)

	return testcontainers.StartGenericContainerRequest{
		Request: testcontainers.ContainerRequest{
			Name:         "mysql",
			Image:        "mysql:8",
			ExposedPorts: []string{":3306"},
			Env: map[string]string{
				"LC_ALL":              "C.UTF-8",
				"MYSQL_DATABASE":      dbName,
				"MYSQL_USER":          dbUser,
				"MYSQL_PASSWORD":      dbPassword,
				"MYSQL_ROOT_PASSWORD": dbPassword,
			},
			WaitingFor: wait.ForAll(
				wait.ForLog("init process done. Ready for start up").
					WithStartupTimeout(5*time.Minute),
				extrawait.ForHealthCheckCmd("mysqladmin", "ping", "-h", "localhost").
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
	return db.RunMigrations(migrationSource, fmt.Sprintf("mysql://%s", DSN(`$MYSQL_DATABASE`, `$MYSQL_USER`, `$MYSQL_PASSWORD`)))
}

// DSN returns the database dsn for connecting to server.
func DSN(dbName, dbUser, dbPassword string) string {
	return fmt.Sprintf("%s:%s@tcp($MYSQL_3306_HOST:$MYSQL_3306_PORT)/%s?charset=utf8&parseTime=true", dbUser, dbPassword, dbName)
}
