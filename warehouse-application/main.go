package main

import (
	"fmt"
	//"warehouse-application/postgresSQL"
	"warehouse-application/postgresSQL"
	"warehouse-application/server"
)

var (
	user     = "postgres"
	password = "ukino"
	host     = "localhost"
	dbname   = "warehouseApp"
	sslmode  = "disable"
)

var DbInfo = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", user, password, host, dbname, sslmode) //DataSourceName

func main() {
	db := postgresSQL.Connect(DbInfo)
	defer postgresSQL.Disconnect(db)
	server.StartWebServer(db)
}
