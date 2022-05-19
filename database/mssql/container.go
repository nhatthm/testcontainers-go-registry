package mssql

import (
	"context"

	"github.com/nhatthm/testcontainers-go-extra"
)

// StartGenericContainer starts a new mssql container.
func StartGenericContainer(ctx context.Context, dbName, dbPassword string, opts ...testcontainers.GenericContainerOption) (testcontainers.Container, error) {
	r := Request(dbName, dbPassword, opts...)

	return testcontainers.StartGenericContainer(ctx, r.Request, r.Options...)
}
