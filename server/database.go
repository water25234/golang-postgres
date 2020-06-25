package server

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	core "github.com/water25234/golang-postgres/core/server"
)

func initDataBase() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		core.GetAppConfig().dbHost,
		core.GetAppConfig().dbPort,
		core.GetAppConfig().dbUsername,
		core.GetAppConfig().dbPassword,
		core.GetAppConfig().dbDatabase)

	db, err := sql.Open(core.GetAppConfig().dbConnection, psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

}
