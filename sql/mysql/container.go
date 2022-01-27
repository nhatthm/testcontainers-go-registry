package testcontainersmysql

import (
	"context"

	"github.com/nhatthm/testcontainers-go-extra"
)

// StartGenericContainer starts a new mysql container.
func StartGenericContainer(ctx context.Context, dbName, dbUser, dbPassword string, opts ...testcontainers.GenericContainerOption) (testcontainers.Container, error) {
	r := Request(dbName, dbUser, dbPassword, opts...)

	return testcontainers.StartGenericContainer(ctx, r.Request, r.Options...)
}
