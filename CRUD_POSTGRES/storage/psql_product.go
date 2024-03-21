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
	psqlCreateProduct  = `INSERT INTO products(name, observations, price, created_at) VALUES($1, $2, $3, $4) RETURNING id`
	psqlGetAllProducts = `SELECT id, name, observations, price, created_at, updated_at FROM products`
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

func (p *PsqlProduct) GetAll() (product.Models, error) {
	// preparamos la sentencia sql
	stmt, err := p.db.Prepare(psqlGetAllProducts)
	if err != nil {
		return nil, err
	}
	// cerramos la conexion si finializa la funcion
	defer stmt.Close()
	// Hacemos la consulta
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	// volvemos a cerrar la conexion si la funcion termina
	defer stmt.Close()
	// creamos el slice de productos
	products := make(product.Models, 0)
	// iteramos por cada registro de la consulta
	for rows.Next() {
		// creamos el producto que nos va a servir para almacenar los campos del registro en cada iteracion
		product := &product.Model{}

		observationNull := sql.NullString{}
		UpdatedAtNull := sql.NullTime{}
		// Mappeamos los datos obtenidos en cada producto y utilizando estructuras intermedias para trabajar con datos nullos
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&observationNull,
			&product.Price,
			&product.CreatedAt,
			&UpdatedAtNull,
		)
		if err != nil {
			return nil, err
		}
		// asignamos valores por defecto al los campos que sean nullos antes de añadirlos al slice
		product.Observation = observationNull.String
		product.UpdatedAt = UpdatedAtNull.Time
		// añadimos el producto al slice de productos
		products = append(products, product)
	}
	// verificamos si salimos del ciclo
	if err := rows.Err(); err != nil {
		return nil, err
	}
	// retornamos el slice de products y nil
	return products, nil
}
