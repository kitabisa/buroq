# go-bootstrap
this is a bootstrap service build with GO for our cookie cutter.

just build it and run. `localhost:[port]/health_check` will give you HTTP status `200`.

## DB driver
    * [mysql] (https://github.com/go-sql-driver/mysql)
    * [postgres] (https://github.com/lib/pq)
    * with [gorp] (https://github.com/go-gorp/gorp) as mapper, why? see benchmark [here] (https://github.com/volatiletech/sqlboiler/blob/master/README.md)

## Cache driver
    * [redis] (https://redis.io)
    * driver using redigo (https://github.com/gomodule/redigo)

## Migration
    * using [sql-migrate] (https://github.com/rubenv/sql-migrate)
    * how to fill the query for migration see on `migrations/yyyymmddhhmmss_migrate_file.sql`
    * command: `migrate` for migrate up, `migratedown` for migrate down

## Router
    * using [go-chi] (https://github.com/go-chi/chi)
    * why? simplicity and easy to read

## Metric
    * using prometheus
    * also running side by side with the api server, with its own port on path `/metrics`

## Others
    * using [perkakas] (https://github.com/kitabisa/perkakas) for any our standard middlewares and for writing response
    * overall using the [golang layout] (https://github.com/golang-standards/project-layout ), with minor changes for own needs

## IMPORTANT!
    * whenever you start new service, please make the documentation using open api specification and put it on folder `/api`