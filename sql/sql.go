package testcontainersql

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Default migration source.
	"github.com/nhatthm/testcontainers-go-extra"
)

// RunMigrations runs database migration when container is ready.
func RunMigrations(sourceURL, databaseURL string) testcontainers.ContainerCallback {
	return func(_ context.Context, _ testcontainers.Container, r testcontainers.ContainerRequest) error {
		m, err := migrate.New(sourceURL, expandDSN(databaseURL, r.Env))
		if err != nil {
			return fmt.Errorf("could not init migration: %w", err)
		}

		if err := m.Up(); err != nil {
			return fmt.Errorf("could not run migration: %w", err)
		}

		return nil
	}
}

func expandDSN(dsn string, env map[string]string) string {
	if l := len(env); l > 0 {
		vars := make([]string, 0, len(env)*2)

		for k, v := range env {
			vars = append(vars, "$"+k, v)
		}

		r := strings.NewReplacer(vars...)

		dsn = r.Replace(dsn)
	}

	return os.ExpandEnv(dsn)
}
