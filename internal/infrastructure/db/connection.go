package db

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	DB *sqlx.DB
}

var dbConn = &DB{}

func OpenConn() (*DB, error) {
	sc := "database.db"

	conn, err := sqlx.Open("sqlite3", sc)

	if err != nil {
		fmt.Println("Erro ao abrir o banco de dados:", err)
		panic(err)
	}

	err = conn.Ping()
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		panic(err)
	}

	dbConn.DB = conn

	fmt.Println("Conexão com o banco de dados SQLite estabelecida com sucesso!")
	return dbConn, err
}

func testDB(conn *sql.DB) error {
	err := conn.Ping()
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		return err
	}
	fmt.Println("Conexão com o banco de dados SQLite estabelecida com sucesso!")
	return nil
}
