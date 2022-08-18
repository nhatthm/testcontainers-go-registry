package mongo

import (
	"context"

	"github.com/nhatthm/testcontainers-go-extra"
)

// StartGenericContainer starts a new mysql container.
//
// Deprecated: Use go.nhat.io/testcontainers-go-registry/mongo.StartGenericContainer instead.
func StartGenericContainer(ctx context.Context, opts ...testcontainers.GenericContainerOption) (testcontainers.Container, error) {
	r := Request(opts...)

	return testcontainers.StartGenericContainer(ctx, r.Request, r.Options...)
}
