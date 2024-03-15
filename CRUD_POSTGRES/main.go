package main

import (
	"log"

	"github.com/TeenBanner/db-go/pkg/product"
	"github.com/TeenBanner/db-go/storage"
	_ "github.com/lib/pq"
)

func main() {
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}
}
