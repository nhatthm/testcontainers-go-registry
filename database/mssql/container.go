package mssql

import (
	"context"

	"go.nhat.io/testcontainers-extra"
)

// StartGenericContainer starts a new mssql container.
//
// Deprecated: Use go.nhat.io/testcontainers-registry/mssql.StartGenericContainer instead.
func StartGenericContainer(ctx context.Context, dbName, dbPassword string, opts ...testcontainers.GenericContainerOption) (testcontainers.Container, error) {
	r := Request(dbName, dbPassword, opts...)

	return testcontainers.StartGenericContainer(ctx, r.Request, r.Options...)
}
