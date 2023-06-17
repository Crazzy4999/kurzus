package db

import (
	"database/sql"
	"fmt"
	config "hangryAPI/configs"

	_ "github.com/lib/pq"
)

var dbSingleton *sql.DB = nil

func GetDB() *sql.DB {
	cfg := config.NewConfig()

	if dbSingleton == nil {
		connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseName, cfg.DatabaseSSLMode)

		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			panic(err)
		}

		dbSingleton = db
	}

	return dbSingleton
}
