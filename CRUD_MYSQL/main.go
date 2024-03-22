package main

import (
	"fmt"

	"github.com/TeenBanner/bases-de-datos-con-go/CRUD_MYSQL/pkg/product"
	"github.com/TeenBanner/bases-de-datos-con-go/CRUD_MYSQL/storage"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	storage.NewMySqlDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	if err := serviceProduct.Migrate(); err != nil {
		fmt.Errorf("Error al migrar la base de datos. Err: %v", err)
	}

}
