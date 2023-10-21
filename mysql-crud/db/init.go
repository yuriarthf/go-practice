package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/spf13/viper"
)

var db *sql.DB

func execInitScript() {
	f, err := os.Open("init.sql")
	defer f.Close()
	if err != nil {
		log.Fatal("Failed to open SQL init script file")
	}
	script := make([]byte, 0, 50)
	f.Read(script)

	_, err = db.Exec(string(script))
	if err != nil {
		log.Fatal("Failed to execute SQL init script")
	}
}

func ConfigMySQL() {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			viper.Get("DB_USER"),
			viper.Get("DB_PASSWORD"),
			viper.Get("DB_HOST"),
			viper.Get("DB_PORT"),
			viper.Get("DB_NAME"),
		),
	)
	if err != nil {
		log.Fatal("Failed to load MySQL DB")
	}

	if err = db.Ping(); err != nil {
		log.Fatal("MySQL DB failed to respond")
	}

	execInitScript()
}

func GetDB() *sql.DB {
	return db
}

func Close() error {
	return db.Close()
}
