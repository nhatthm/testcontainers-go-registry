package elasticsearch_test

import (
	"context"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql" // Database driver
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.nhat.io/testcontainers-go-extra"

	es "go.nhat.io/testcontainers-go-registry/elasticsearch"
)

func TestRunMigrations(t *testing.T) {
	t.Parallel()

	c, err := es.StartGenericContainer(context.Background(),
		testcontainers.ContainerCallback(func(context.Context, testcontainers.Container, testcontainers.ContainerRequest) error {
			esURL := os.ExpandEnv(es.URL())

			request, err := http.NewRequest(http.MethodGet, esURL+"_cat/indices?format=json", nil)
			if err != nil {
				return err
			}

			resp, err := http.DefaultClient.Do(request)
			if err != nil {
				return err
			}

			defer resp.Body.Close() // nolint: errcheck

			data, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			assert.True(t, strings.HasPrefix(string(data), "["))

			return nil
		}),
	)

	assert.NoError(t, err)
	require.NotNil(t, c)

	_ = testcontainers.StopGenericContainers(context.Background(), c) // nolint: errcheck
}
