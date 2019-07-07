# Golang Rest API Template - MongoDB, MariaDB

## Requires

`Go 1.12.x`
`Docker`
`Docker Compose`

## Begin

```bash
make deps
```

## Build

```bash
make all
```

## Run

```bash
make run-deps
make run
make stop-deps
```

Verify in <http://localhost:8010/>

## Run integration tests

```bash
make itest
```

## Pedendencies

GoConvey - Testing tool

```bash
dep ensure -add github.com/smartystreets/goconvey
```

MariaDB driver

```bash
dep ensure -add github.com/go-sql-driver/mysql
```

MongoDB driver

```bash
dep ensure -add go.mongodb.org/mongo-driver/mongo@~1.0.0
```

Google uuid

```bash
dep ensure -add github.com/google/uuid
```

Go yaml

```bash
dep ensure -add  github.com/go-yaml/yaml
```
