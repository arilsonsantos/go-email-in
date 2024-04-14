package db

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	//_ "github.com/mattn/go-sqlite3"
	//_ "github.com/jackc/pgx/v5"
)

type DB struct {
	DB *sqlx.DB
}

var dbConn = &DB{}

func OpenConn() (*DB, error) {
	//sc := "database.db"
	strDb := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	conn, err := sqlx.Open("postgres", strDb)

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
