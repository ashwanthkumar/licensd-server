package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var database *sql.DB

func initDatabase() {
	dbType := viper.GetString("database.type")
	dbDSN := viper.GetString("database.dsn")
	db, err := sql.Open(dbType, dbDSN)
	if err != nil {
		panic(fmt.Errorf("Fatal error db init failed: %s \n", err))
	}
	database = db
}
