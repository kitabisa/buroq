# go-bootstrap
this is a bootstrap service build with GO for our cookie cutter

## DB driver
    * [mysql] (https://github.com/go-sql-driver/mysql)
    * [postgre] (https://github.com/lib/pq)
    * with [gorp] (https://github.com/go-gorp/gorp) as mapper, why? see benchmark [here] (https://github.com/volatiletech/sqlboiler/blob/master/README.md)

## Cache driver
    * [redis] (https://redis.io)
    * driver using redigo (https://github.com/gomodule/redigo)