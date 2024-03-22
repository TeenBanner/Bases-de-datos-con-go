package storage

import (
	"database/sql"
	"fmt"
)

const (
	psqlMigrateProduct = `
	CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL PRIMARY KEY,
		name VARCHAR(50) NOT NULL,
		observations VARCHAR(100),
		Created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		Updated_at TIMESTAMP
	);
	`
)

type psqlProduct struct {
	Db *sql.DB
}

func NewPsqlProduct(db *sql.DB) *psqlProduct {
	return &psqlProduct{
		Db: db,
	}
}

func (p *psqlProduct) Migrate() error {
	stmt, err := p.Db.Prepare(psqlMigrateProduct)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migracion de producto realizada satisfactoriamente")

	return nil
}
