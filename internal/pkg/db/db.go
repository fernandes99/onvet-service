package db

import (
	"database/sql"
	"fmt"
	"log"
	"onvet/internal/app/pkg/configuration"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB     *sql.DB
	config *configuration.Config
)

func StartConnection() {
	config = configuration.Load() // TODO: shared all packages

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	var err error
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Println("failed to open database connection: ", err)
		panic(err)
	}

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxLifetime(5 * time.Minute)

	fmt.Println("database connected!")
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := DB.Exec(query, args...)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := DB.Query(query, args...)

	if err != nil {
		return nil, err
	}

	return rows, nil
}

func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
