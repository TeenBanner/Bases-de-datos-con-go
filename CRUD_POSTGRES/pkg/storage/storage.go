package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewPostgresDB() {
	once.Do(func() {
		var err error
		var connStr string = "postgres://edteam:edteam@localhost:3000/godb?sslmode=disable"
		// aqui se ejecuta una sola vez
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Can't ping db: %v", err)
		}

		fmt.Println("Conectado a postgres")
	})
}

// pool returns a unique instace of db
func Pool() *sql.DB {
	return db
}
