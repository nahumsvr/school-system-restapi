package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "host=localhost user=nahum password=nahumpassword dbname=gorm port=5432"
var DB *gorm.DB
var err error

func Connect() {
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error al conectar a la base de datos:", err)
		return
	}
	fmt.Println("Conexión exitosa a la base de datos")
}
