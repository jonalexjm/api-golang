package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"api-rest-go/internal/service"
	"api-rest-go/internal/store"
	"api-rest-go/internal/transport"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("DELETE /books/{id}    - Eliminar un libro por ID")
	//conectar con sqllite
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//crear tabla si no existe

	q := `CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		author TEXT NOT NULL
	);`

	if _, err := db.Exec(q); err != nil {
		log.Fatal(err.Error())
	}

	//inyectar dependencias
	booksStore := store.New(db)
	booksService := service.New(booksStore)
	booksHandler := transport.New(booksService)
	//configurar rutas

	http.HandleFunc("/books", booksHandler.HandleBooks)
	http.HandleFunc("/books/", booksHandler.HandleBookByID)

	fmt.Println("servidor escuchando en :8080")
	fmt.Println("API Endpoint")
	fmt.Println("GET    /books         - Obtener todos los libros")
	fmt.Println("POST   /books         - Crear un nuevo libro")
	fmt.Println("GET    /books/{id}    - Obtener un libro por ID")
	fmt.Println("PUT    /books/{id}    - Actualizar un libro por ID")
	fmt.Println("DELETE /books/{id}    - Eliminar un libro por ID")

	// escuchar peticiones http
	log.Fatal(http.ListenAndServe(":8080", nil))
}
