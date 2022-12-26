module go.nhat.io/testcontainers-registry/mysql

go 1.18

require (
	github.com/go-sql-driver/mysql v1.7.0
	github.com/golang-migrate/migrate/v4 v4.15.2
	github.com/stretchr/testify v1.8.1
	github.com/testcontainers/testcontainers-go v0.17.0
	go.nhat.io/testcontainers-extra v0.8.0
	go.nhat.io/testcontainers-registry v0.10.0
)

require (
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/Microsoft/go-winio v0.6.0 // indirect
	github.com/cenkalti/backoff/v4 v4.2.0 // indirect
	github.com/containerd/containerd v1.6.12 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/docker/distribution v2.8.1+incompatible // indirect
	github.com/docker/docker v20.10.21+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/klauspost/compress v1.15.13 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/moby/patternmatcher v0.5.0 // indirect
	github.com/moby/sys/sequential v0.5.0 // indirect
	github.com/moby/term v0.0.0-20221205130635-1aeaba878587 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0-rc2 // indirect
	github.com/opencontainers/runc v1.1.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	golang.org/x/mod v0.7.0 // indirect
	golang.org/x/net v0.4.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/tools v0.4.0 // indirect
	google.golang.org/genproto v0.0.0-20221207170731-23e4bf6bdc37 // indirect
	google.golang.org/grpc v1.51.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gotest.tools/v3 v3.4.0 // indirect
)

replace (
	// Given the version includes the Compose dependency, and the Docker folks added a replace directive until the upcoming
	// Docker 22.06 release is out, we were forced to add it too, causing consumers of Testcontainers for Go to add the
	// following replace directive to their go.mod files. We expect this to be removed in the next releases of
	// Testcontainers for Go.
	github.com/docker/cli => github.com/docker/cli v20.10.3-0.20221013132413-1d6c6e2367e2+incompatible
	github.com/docker/docker => github.com/docker/docker v20.10.3-0.20221013203545-33ab36d6b304+incompatible
	github.com/moby/buildkit => github.com/moby/buildkit v0.10.1-0.20220816171719-55ba9d14360a

	// For k8s dependencies, we use a replace directive, to prevent them being
	// upgraded to the version specified in containerd, which is not relevant to the
	// version needed.
	// See https://github.com/docker/buildx/pull/948 for details.
	// https://github.com/docker/buildx/blob/v0.8.1/go.mod#L62-L64
	k8s.io/api => k8s.io/api v0.22.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.22.4
	k8s.io/client-go => k8s.io/client-go v0.22.4
)
