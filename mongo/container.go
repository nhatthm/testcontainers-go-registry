package mongo

import (
	"context"

	"go.nhat.io/testcontainers-extra"
)

// StartGenericContainer starts a new mysql container.
func StartGenericContainer(ctx context.Context, opts ...testcontainers.GenericContainerOption) (testcontainers.Container, error) {
	r := Request(opts...)

	return testcontainers.StartGenericContainer(ctx, r.Request, r.Options...)
}
