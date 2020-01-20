[![Go Report Card](https://goreportcard.com/badge/github.com/kitabisa/buroq)](https://goreportcard.com/report/github.com/kitabisa/buroq)
![technology:go](https://img.shields.io/badge/technology-go-blue.svg)
[![Maintainability](https://api.codeclimate.com/v1/badges/b6e84c6b2bb2a819b198/maintainability)](https://codeclimate.com/github/kitabisa/buroq/maintainability)
[![CircleCI](https://circleci.com/gh/kitabisa/buroq.svg?style=svg)](https://circleci.com/gh/kitabisa/buroq)

# buroq

this is a bootstrap service build with GO for our cookie cutter.

just build it and run. `localhost:[port]/health_check` will give you HTTP status `200`.

## DB driver

* [mysql](https://github.com/go-sql-driver/mysql)
* [postgres](https://github.com/lib/pq)
* with [gorp](https://github.com/go-gorp/gorp) as mapper, why? see benchmark [here](https://github.com/volatiletech/sqlboiler/blob/master/README.md)

## Cache driver

* [redis](https://redis.io)
* driver using [redigo](https://github.com/gomodule/redigo)

## Migration
Before you do the migration, please specify what DB that you want to migrate in configuration file with key `is_migration_enable`.
* using [sql-migrate](https://github.com/rubenv/sql-migrate)
* command: `migrate` for migrate up
* command: `migratedown` for migrate down
* command: `migratenew [migration name]` for create migration file

## Seeding (Work in progress)

* seeds for starter can be put on `migrations/seeds/{number}.{table-name}.sql`

## Router

* using [go-chi](https://github.com/go-chi/chi)
* why? simplicity and easy to read

## Metric (Work in progress)

* using [InfluxDB](https://www.influxdata.com)

## GraphQL

* you can see some docs about graphql integration with buroq under `internal/app/graphql`

## Others

* using [perkakas](https://github.com/kitabisa/perkakas) for any our standard middlewares and for writing response
* overall using the [golang layout](https://github.com/golang-standards/project-layout), with minor changes for own needs

---

## IMPORTANT

* whenever you start new service, please make the documentation using open api specification and put it on folder `/api`

## Author

* Arditya Wahyu N - *initial work* - [ardityawahyu](https://github.com/ardityawahyu)
