package mssql

import (
	"context"

	"github.com/nhatthm/testcontainers-go-extra"
)

// StartGenericContainer starts a new mssql container.
//
// Deprecated: Use go.nhat.io/testcontainers-go-registry/mssql.StartGenericContainer instead.
func StartGenericContainer(ctx context.Context, dbName, dbPassword string, opts ...testcontainers.GenericContainerOption) (testcontainers.Container, error) {
	r := Request(dbName, dbPassword, opts...)

	return testcontainers.StartGenericContainer(ctx, r.Request, r.Options...)
}
