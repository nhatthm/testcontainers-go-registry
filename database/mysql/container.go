package mysql

import (
	"context"

	"go.nhat.io/testcontainers-go-extra"
)

// StartGenericContainer starts a new mysql container.
//
// Deprecated: Use go.nhat.io/testcontainers-go-registry/mysql.StartGenericContainer instead.
func StartGenericContainer(ctx context.Context, dbName, dbUser, dbPassword string, opts ...testcontainers.GenericContainerOption) (testcontainers.Container, error) {
	r := Request(dbName, dbUser, dbPassword, opts...)

	return testcontainers.StartGenericContainer(ctx, r.Request, r.Options...)
}
