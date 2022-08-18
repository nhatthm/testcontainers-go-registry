# Testcontainers-Go Common Image Registry

[![GitHub Releases](https://img.shields.io/github/v/release/nhatthm/testcontainers-go-registry)](https://github.com/nhatthm/testcontainers-go-registry/releases/latest)
[![Build Status](https://github.com/nhatthm/testcontainers-go-registry/actions/workflows/test.yaml/badge.svg)](https://github.com/nhatthm/testcontainers-go-registry/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/nhatthm/testcontainers-go-registry/branch/master/graph/badge.svg?token=eTdAgDE2vR)](https://codecov.io/gh/nhatthm/testcontainers-go-registry)
[![Go Report Card](https://goreportcard.com/badge/github.com/nhatthm/testcontainers-go-registry)](https://goreportcard.com/report/github.com/nhatthm/testcontainers-go-registry)
[![GoDevDoc](https://img.shields.io/badge/dev-doc-00ADD8?logo=go)](https://pkg.go.dev/github.com/nhatthm/testcontainers-go-registry)
[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.com/donate/?hosted_button_id=PJZSGJN57TDJY)

Common Image Registry for Testcontainers-Go

## Prerequisites

- `Go >= 1.17`

## Install

```bash
go get go.nhat.io/testcontainers-go-registry
```

## Mongo

```go
package example

import (
	"context"

	"go.nhat.io/testcontainers-go-registry/mongo"
	testcontainers "github.com/nhatthm/testcontainers-go-extra"
)

const (
	dbName          = "test"
	migrationSource = "file://./resources/migrations/"
)

func startMongoDB() (testcontainers.Container, error) {
	return mongo.StartGenericContainer(context.Background(),
		mongo.RunMigrations(migrationSource, dbName),
	)
}
```

## MySQL

```go
package example

import (
	"context"

	"go.nhat.io/testcontainers-go-registry/mysql"
	testcontainers "github.com/nhatthm/testcontainers-go-extra"
)

const (
	dbName          = "test"
	dbUser          = "test"
	dbPassword      = "test"
	migrationSource = "file://./resources/migrations/"
)

func startMySQL() (testcontainers.Container, error) {
	return mysql.StartGenericContainer(context.Background(),
		dbName, dbUser, dbPassword,
		mysql.RunMigrations(migrationSource),
	)
}
```

## Postgres

```go
package example

import (
	"context"

	"go.nhat.io/testcontainers-go-registry/postgres"
	testcontainers "github.com/nhatthm/testcontainers-go-extra"
)

const (
	dbName          = "test"
	dbUser          = "test"
	dbPassword      = "test"
	migrationSource = "file://./resources/migrations/"
)

func startPostgres() (testcontainers.Container, error) {
	return postgres.StartGenericContainer(context.Background(),
		dbName, dbUser, dbPassword,
		postgres.RunMigrations(migrationSource),
	)
}
```

## SQL Server

```go
package example

import (
	"context"

	testcontainers "github.com/nhatthm/testcontainers-go-extra"
	"go.nhat.io/testcontainers-go-registry/mssql"
)

const (
	dbName          = "test"
	dbPassword      = "My!StrongPassword"
	migrationSource = "file://./resources/migrations/"
)

func startSQLServer() (testcontainers.Container, error) {
	return mssql.StartGenericContainer(context.Background(),
		dbName, dbPassword,
		mssql.RunMigrations(migrationSource),
	)
}
```

## Options

### Change Image Tag

```go
package example

import (
	"context"

	"go.nhat.io/testcontainers-go-registry/postgres"
	testcontainers "github.com/nhatthm/testcontainers-go-extra"
)

const (
	dbName     = "test"
	dbUser     = "test"
	dbPassword = "test"
)

func startPostgres() (testcontainers.Container, error) {
	return postgres.StartGenericContainer(context.Background(),
		dbName, dbUser, dbPassword,
		testcontainers.WithImageTag("13-alpine"),
	)
}
```

### Change Image Name

```go
package example

import (
	"context"

	testcontainers "github.com/nhatthm/testcontainers-go-extra"
	"go.nhat.io/testcontainers-go-registry/mysql"
)

const (
	dbName     = "test"
	dbUser     = "test"
	dbPassword = "test"
)

func startMySQL() (testcontainers.Container, error) {
	return mysql.StartGenericContainer(context.Background(),
		dbName, dbUser, dbPassword,
		testcontainers.WithImageName("mariadb"),
		testcontainers.WithImageTag("10.7"),
	)
}
```

## Donation

If this project help you reduce time to develop, you can give me a cup of coffee :)

### Paypal donation

[![paypal](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/donate/?hosted_button_id=PJZSGJN57TDJY)

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;or scan this

<img src="https://user-images.githubusercontent.com/1154587/113494222-ad8cb200-94e6-11eb-9ef3-eb883ada222a.png" width="147px" />
