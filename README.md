# Golang REST API with MongoDB, MariaDB, Prometheus

## Requires

`Go 1.18.x`
`Make`
`Docker`
`Docker Compose`

## Install

```bash
make install
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

## Curls

Ping
```bash
curl --location --request GET 'http://localhost:8010/ping'
```

Monitors
```bash
curl --location --request GET 'http://localhost:8010/monitor'
```

Metrics (Prometheus)
```bash
curl --location --request GET 'http://localhost:8010/metrics'
```