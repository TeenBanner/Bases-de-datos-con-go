package main

import (
	"github.com/TeenBanner/db-go/pkg/storage"
	_ "github.com/lib/pq"
)

func main() {
	storage.NewPostgresDB()
	storage.NewPostgresDB()
	storage.NewPostgresDB()

}
