package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewMySqlDB() {
	once.Do(func() {
		var err error
		var connStr string = "root:edteam@tcp(localhost:8000)/MySqlGoDB?charset=utf8mb4&parseTime=True&loc=Local"

		db, err = sql.Open("mysql", connStr)
		if err != nil {
			log.Fatalf("No se pudo conectar a la Base de datos. Error: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Cannot Ping DB. ERR: %v", err)
		}

		fmt.Println("Conectado a mysql")
	})
}

func Pool() *sql.DB {
	return db
}
