package config

import (
	"database/sql"
	"fmt"
	"log"
)

type App struct {
	DB *sql.DB
}

//type Workers struct {
//	Workers int
//}
func (a *App) InitializeDb(user, password, dbname, sslmode string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", user, password, dbname, sslmode)

	var err error
	if a.DB == nil {
		a.DB, err = sql.Open("postgres", connectionString)
		if err != nil {
			log.Fatal(err)
		}
	}
}

//var InitWorkers = Workers{Workers: 1}