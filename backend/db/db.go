package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func Connect(connStr string) {
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error abriendo la base de datos: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	log.Println("Conexión a la base de datos establecida.")
}
