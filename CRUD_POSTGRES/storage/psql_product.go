package storage

import (
	"database/sql"
	"fmt"

	"github.com/TeenBanner/db-go/pkg/product"
)

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT product_id_pk PRIMARY KEY (id)
	)`
	psqlCreateProduct = `INSERT INTO products(name, observations, price, created_at) VALUES($1, $2, $3, $4) RETURNING id`
)

// Psql used for work with postgres -product
type PsqlProduct struct {
	db *sql.DB
}

// NewsqlProduct
func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

func (p *PsqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migracion de producto ejecutada correctamente")

	return nil
}

// method create Insert a product into database implemnet the interface product.Storage
func (p *PsqlProduct) Create(m *product.Model) error {
	// Preparamos la sentencia sql con la constante correspondiente
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	// Ejecutamos la consulta y le pasamos los campos que estamos indicando con los marcadores de poscicion tambien verificando si es un valor nullo el que estamos insertando en la DB
	err = stmt.QueryRow(m.Name, stringtoNull(m.Observation), m.Price, m.CreatedAt).Scan(&m.ID) // Escaneamos el valor retornado de la consulta y lo mappeamos en el campo id del modelo que recibimos usando la direccion de memoria del campo ID del modelo recivido
	if err != nil {
		return err
	}

	fmt.Println("Se creo el producto correctamente")

	return nil
}
