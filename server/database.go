package server

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	core "github.com/water25234/golang-postgres/core/server"
)

var db *sql.DB

func InitDataBase() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		core.GetAppConfig().DBHost,
		core.GetAppConfig().DBPort,
		core.GetAppConfig().DBUsername,
		core.GetAppConfig().DBPassword,
		core.GetAppConfig().DBDatabase)

	db, err := sql.Open(core.GetAppConfig().DBConnection, psqlInfo)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
