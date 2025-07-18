package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/AlbMP96/backend/db"
	"github.com/AlbMP96/backend/handlers"

	_ "github.com/lib/pq"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola desde Go!")
}

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db.Connect(psqlInfo)

	http.HandleFunc("/", handler)
	http.HandleFunc("/api/users/create", handlers.CreateUserHandler)
	http.HandleFunc("/api/users", handlers.GetUserHandler)

	fmt.Println("Escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", nil)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Println("Cerrando conexión a la base de datos...")
	db.DB.Close()
}
