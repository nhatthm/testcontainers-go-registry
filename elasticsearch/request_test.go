package elasticsearch_test

import (
	"context"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.nhat.io/testcontainers-extra"

	es "go.nhat.io/testcontainers-registry/elasticsearch"
)

func TestRunMigrations(t *testing.T) {
	t.Parallel()

	c, err := es.StartGenericContainer(context.Background(),
		testcontainers.ContainerCallback(func(context.Context, testcontainers.Container, testcontainers.ContainerRequest) error {
			esURL := os.ExpandEnv(es.URL())

			request, err := http.NewRequest(http.MethodGet, esURL+"_cat/indices?format=json", nil)
			if err != nil {
				return err //nolint: wrapcheck
			}

			resp, err := http.DefaultClient.Do(request)
			if err != nil {
				return err //nolint: wrapcheck
			}

			defer resp.Body.Close() // nolint: errcheck,gosec

			data, err := io.ReadAll(resp.Body)
			if err != nil {
				return err //nolint: wrapcheck
			}

			assert.True(t, strings.HasPrefix(string(data), "["))

			return nil
		}),
	)

	assert.NoError(t, err)
	require.NotNil(t, c)

	_ = testcontainers.StopGenericContainers(context.Background(), c) // nolint: errcheck
}
