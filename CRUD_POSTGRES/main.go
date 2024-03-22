package main

import (
	"log"

	"github.com/TeenBanner/db-go/pkg/invoice"
	invoiceheader "github.com/TeenBanner/db-go/pkg/invoiceHeader"
	invoiceitem "github.com/TeenBanner/db-go/pkg/invoiceItem"
	"github.com/TeenBanner/db-go/storage"
	_ "github.com/lib/pq"
)

func main() {
	// hace la conexion con la BD
	storage.NewPostgresDB()
	// migraciones
	// storageProduct := storage.NewPsqlProduct(storage.Pool())
	// serviceProduct := product.NewService(storageProduct)

	// if err := serviceProduct.Migrate(); err != nil {
	// 	log.Fatalf("product.Migrate: %v", err)
	// }

	// storageInvoiceHeader := storage.NewpsqlInvoiceHeader(storage.Pool())
	// serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)

	// if err := serviceInvoiceHeader.Migrate(); err != nil {
	// 	log.Fatalf("serviceInvoiceHeader.Migrate: %v", err)
	// }

	// StorageInvoiceItem := storage.NewpsqlInvoiceItem(storage.Pool())
	// ServiceInvoiceItem := invoiceitem.NewService(StorageInvoiceItem)

	// if err := ServiceInvoiceItem.Migrate(); err != nil {
	// 	log.Fatalf("InvoiceItemError: %v", err)
	// }
	// Create Method using storageproduct && service product

	// instaciamos el producto a crear
	/*m := &product.Model{
		Name:        "Bases de datos con Go",
		Price:       46,
		Observation: "on fire",
	}
	if err := serviceProduct.Create(m); err != nil {
		log.Fatalf("Product Create at main: %v", err)
	}
	// revisamos el id del producto insertado y la fecha de creacion que se creo solo con llamar al metodo migrate

	fmt.Printf("%+v", m)
	*/

	// GetALl Method

	// products, err := serviceProduct.GetAll()
	// if err != nil {
	// 	log.Fatalf("Product.GetAll(): %v", err)
	// }

	// fmt.Println(products)

	// fmt.Println("Get By ID method: ")

	// product, err := serviceProduct.GetByID(0)
	// switch {
	// case errors.Is(err, sql.ErrNoRows):
	// 	fmt.Println("El id no existe")
	// case err != nil:
	// 	log.Fatalf("Error: %v", err)
	// default:
	// 	fmt.Println(product)
	// }

	// Update method

	// productToUpdate := &product.Model{
	// 	ID:    40,
	// 	Name:  "DB con go",
	// 	Price: 80,
	// }

	// err := serviceProduct.Update(productToUpdate)
	// if err != nil {
	// 	log.Fatalf("Error al actualizar el producto %v", err)
	// }
	// delete method

	// err := serviceProduct.Delete(3)
	// if err != nil {
	// 	log.Fatalf("Error al eliminar el producto: %v", err)
	// }

	// transaction method
	storageInvoiceHeader := storage.NewpsqlInvoiceHeader(storage.Pool())
	storageinvoiceItem := storage.NewpsqlInvoiceItem(storage.Pool())
	storageinvoice := storage.NewPsqlInvoice(
		storage.Pool(),
		storageInvoiceHeader,
		storageinvoiceItem,
	)

	factura := &invoice.Model{
		Header: &invoiceheader.Model{
			Client: "Alvaro Felipe",
		},
		Items: invoiceitem.Models{
			&invoiceitem.Model{ProductID: 1},
			&invoiceitem.Model{ProductID: 99},
		},
	}
	invoiceService := invoice.NewService(storageinvoice)
	if err := invoiceService.Create(factura); err != nil {
		log.Fatalf("Error invoice.Create: %v", err)
	}
}
