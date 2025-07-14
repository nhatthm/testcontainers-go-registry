package mssql

import (
	"context"

	"go.nhat.io/testcontainers-extra"
)

// StartGenericContainer starts a new mssql container.
func StartGenericContainer(ctx context.Context, dbName, dbPassword string, opts ...testcontainers.GenericContainerOption) (testcontainers.Container, error) {
	r := Request(dbName, dbPassword, opts...)

	return testcontainers.StartGenericContainer(ctx, r.Request, r.Options...) //nolint: wrapcheck
}
