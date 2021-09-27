# Golang Rest API Template - MongoDB, MariaDB

## Requires

`Go 1.17.x`
`Make`
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
make deps-start
make run
make deps-stop
```

Verify in <http://localhost:8010/>

## Run integration tests

```bash
make itest
```

## Pedendencies

GoConvey - Testing tool

```bash
github.com/smartystreets/goconvey
```

MariaDB driver

```bash
github.com/go-sql-driver/mysql
```

MongoDB driver

```bash
go.mongodb.org/mongo-driver/mongo@~1.0.0
```

Google uuid

```bash
github.com/google/uuid
```

Go yaml

```bash
github.com/go-yaml/yaml
```
