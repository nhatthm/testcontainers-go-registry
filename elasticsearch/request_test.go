package testcontainerselasticsearch_test

import (
	"context"
	"io"
	"net/http"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql" // Database driver
	"github.com/nhatthm/testcontainers-go-extra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	testcontainerselasticsearch "github.com/nhatthm/testcontainers-go-registry/elasticsearch"
)

func TestRunMigrations(t *testing.T) {
	t.Parallel()

	c, err := testcontainerselasticsearch.StartGenericContainer(context.Background(),
		testcontainers.ContainerCallback(func(context.Context, testcontainers.Container, testcontainers.ContainerRequest) error {
			esURL := os.ExpandEnv(testcontainerselasticsearch.URL())

			request, err := http.NewRequest(http.MethodGet, esURL+"_cat/indices?format=json", nil)
			if err != nil {
				return err
			}

			resp, err := http.DefaultClient.Do(request)
			if err != nil {
				return err
			}

			data, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			assert.Equal(t, `[]`, string(data))

			return nil
		}),
	)

	assert.NoError(t, err)
	require.NotNil(t, c)

	_ = testcontainers.StopGenericContainers(context.Background(), c) // nolint: errcheck
}
