package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

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

// StringToNull help us to work with null values on DataBase
func stringtoNull(s string) sql.NullString { // Recibe un string y devuelve una estructura NullString
	// Create Null Value
	null := sql.NullString{
		String: s,
	}
	// verificamos si nos estan mandando un valor nullo desde la base de datos o si nos estan enviando un valor
	if null.String != "" {
		null.Valid = true
	}
	// retornamos el valor null
	return null
}

// TimeToNull it's a helper to work with null dates
func timeToNull(t time.Time) sql.NullTime { // recive una estructura de tipo Time y devuelve un sql.NullTime
	null := sql.NullTime{Time: t} // Creamos la variable null la cual es de tipo NullTime y le pasamos el valor t como valor del campo time
	if !null.Time.IsZero() {      // verificamos si null no tiene el valor cero para asignar que null no es un dato nullo
		null.Valid = true
	}
	// retornamos null
	return null
}
