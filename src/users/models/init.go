package models

import (
	"database/sql"
	"go_project/src/config"
)

type DbConnect struct {
	db *sql.DB
}

var conn = new(DbConnect)

func Init (a config.App) {
	conn.db = a.DB
}