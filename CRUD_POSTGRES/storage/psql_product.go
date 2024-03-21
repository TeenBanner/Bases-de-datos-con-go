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
	psqlGetProductByID = psqlGetAllProducts + " WHERE id = $1"
	psqlUpdateProduct  = "UPDATE products SET name = $1, observations = $2, price =$3, updated_at = $4 WHERE id = $5"
	psqlDeleteProduct  = "DELETE FROM products WHERE id = $1"
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
		product, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		products = append(products, product)
	}
	// verificamos si salimos del ciclo
	if err := rows.Err(); err != nil {
		return nil, err
	}
	// retornamos el slice de products y nil
	return products, nil
}

type scanner interface {
	Scan(dest ...interface{}) error
}

// GetById implement product.Storage Getting a product by it id
func (p *PsqlProduct) GetByID(ID uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(psqlGetProductByID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return scanRowProduct(stmt.QueryRow(ID))
}

// scanRowProducts it's used for mapp the elements comming from DB INTO a Product Model
func scanRowProduct(s scanner) (*product.Model, error) {
	ModelProduct := &product.Model{}
	// estructuras para manejar nullos
	observationNull := sql.NullString{}
	UpdatedAtNull := sql.NullTime{}
	// Mappeamos los datos obtenidos en cada producto y utilizando estructuras intermedias para trabajar con datos nullos
	err := s.Scan(
		&ModelProduct.ID,
		&ModelProduct.Name,
		&observationNull,
		&ModelProduct.Price,
		&ModelProduct.CreatedAt,
		&UpdatedAtNull,
	)
	if err != nil {
		return &product.Model{}, err
	}
	// asignamos valores por defecto al los campos que sean nullos antes de añadirlos al slice
	ModelProduct.Observation = observationNull.String
	ModelProduct.UpdatedAt = UpdatedAtNull.Time
	// retornamos el scanneo si todo sale bien
	return ModelProduct, nil
}

// Update implements the interface Storaga, it's used for Update a product
func (p *PsqlProduct) Update(m *product.Model) error { // recibe un puntero a un producto  y devuelve un posible error
	// Preparamos la declaracion sql
	stmt, err := p.db.Prepare(psqlUpdateProduct)
	if err != nil {
		return err
	}
	// Cerramos la conexion si finaliza la funcion por un error
	defer stmt.Close()
	// hacemos la ejecucion de la sentencia prepararada y le pasamos los valores de campos a actualizar en los marcadores de posicion con los valores del puntero de producto de los parametros
	res, err := stmt.Exec(
		m.Name,
		stringtoNull(m.Observation),
		m.Price,
		m.UpdatedAt,
		m.ID,
	)
	if err != nil {
		return err
	}
	rowsaffcted, err := res.RowsAffected()

	if rowsaffcted != 1 {
		return fmt.Errorf("No existe un registro con el id: %v", m.ID)
	}
	fmt.Println("Producto actualizado correctamente")

	return nil
}

func (p *PsqlProduct) Delete(Id uint) error {
	stmt, err := p.db.Prepare(psqlDeleteProduct)
	if err != nil {
		return err
	}

	defer db.Close()

	res, err := stmt.Exec(Id)
	if err != nil {
		return err
	}

	rowsAffcted, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffcted != 1 {
		return fmt.Errorf("No existe un registro con el id: %v", Id)
	}

	fmt.Println("Producto Eliminado satisfactoriamente")

	return nil
}
