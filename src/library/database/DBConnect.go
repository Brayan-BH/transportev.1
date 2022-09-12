package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	_"github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
)

//Connection retorna la conexion con la base de datos
func Connection() *sql.DB {
	var db *sql.DB
	err_load := godotenv.Load()
	if err_load != nil {
		log.Fatal("Error al configurar viariable de entorno")
	}

	server := os.Getenv("DB_SERVER")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")

	connStrings := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;",
		server, user, password, database)

	db, err := sql.Open("sqlserver", connStrings)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	ctx := context.Background()
	err = db.PingContext(ctx)

	if err != nil {
		log.Fatal("Error creating connection: ", err.Error())
	}
	// fmt.Printf("Conectado!\n")
	return db
}
