package main

import (
	"fmt"
	"log"

	invoiceheader "github.com/TeenBanner/db-go/pkg/invoiceHeader"
	invoiceitem "github.com/TeenBanner/db-go/pkg/invoiceItem"
	"github.com/TeenBanner/db-go/pkg/product"
	"github.com/TeenBanner/db-go/storage"
	_ "github.com/lib/pq"
)

func main() {
	storage.NewPostgresDB()
	// migraciones
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}

	storageInvoiceHeader := storage.NewpsqlInvoiceHeader(storage.Pool())
	serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)

	if err := serviceInvoiceHeader.Migrate(); err != nil {
		log.Fatalf("serviceInvoiceHeader.Migrate: %v", err)
	}

	StorageInvoiceItem := storage.NewpsqlInvoiceItem(storage.Pool())
	ServiceInvoiceItem := invoiceitem.NewService(StorageInvoiceItem)

	if err := ServiceInvoiceItem.Migrate(); err != nil {
		log.Fatalf("InvoiceItemError: %v", err)
	}
	// Create Method using storageproduct && service product

	// instaciamos el producto a crear
	m := &product.Model{
		Name:        "Bases de datos con Go",
		Price:       46,
		Observation: "on fire",
	}
	if err := serviceProduct.Create(m); err != nil {
		log.Fatalf("Product Create at main: %v", err)
	}
	// revisamos el id del producto insertado y la fecha de creacion que se creo solo con llamar al metodo migrate

	fmt.Printf("%+v", m)

}
