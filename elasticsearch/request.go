package elasticsearch

import (
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-units"
	"go.nhat.io/testcontainers-extra"
	extrawait "go.nhat.io/testcontainers-extra/wait"
)

// Request creates a new request for starting an elasticsearch server.
func Request(opts ...testcontainers.GenericContainerOption) testcontainers.StartGenericContainerRequest {
	finalOpts := make([]testcontainers.GenericContainerOption, 1, len(opts)+1)
	finalOpts[0] = testcontainers.PopulateHostPortEnv
	finalOpts = append(finalOpts, opts...)

	return testcontainers.StartGenericContainerRequest{
		Request: testcontainers.ContainerRequest{
			Name:         "elasticsearch",
			Image:        "docker.elastic.co/elasticsearch/elasticsearch:9.0.0",
			ExposedPorts: []string{":9200"},
			Env: map[string]string{
				"xpack.security.enabled": "false",
				"discovery.type":         "single-node",
				"ES_JAVA_OPTS":           "-Xms512m -Xmx512m",
			},
			HostConfigModifier: func(c *container.HostConfig) {
				c.Resources.Ulimits = []*units.Ulimit{ //nolint: staticcheck
					{
						Name: "memlock",
						Hard: -1,
						Soft: -1,
					},
					{
						Name: "nofile",
						Hard: 65536,
						Soft: 65536,
					},
				}
			},
			WaitingFor: extrawait.ForHealthCheckCmd("curl", "--silent", "--fail", "localhost:9200/_cluster/health").
				WithRetries(3).
				WithStartPeriod(time.Minute).
				WithTestTimeout(5 * time.Second).
				WithTestInterval(10 * time.Second),
		},
		Options: finalOpts,
	}
}

// URL returns the url for connecting to server.
func URL() string {
	return "http://$ELASTICSEARCH_9200_HOST:$ELASTICSEARCH_9200_PORT/"
}
