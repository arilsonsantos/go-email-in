package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strconv"
	"time"
)

type DB struct {
	DB *sql.DB
}

func OpenConn() (*DB, error) {
	var dbConn = &DB{}
	strDB := getStrDB()

	conn, err := sql.Open(driverName, strDB)

	conn.SetMaxIdleConns(getEnvInt("DB_MAX_IDLE_CONN", 5))
	conn.SetMaxOpenConns(getEnvInt("DB_MAX_OPEN_CONN", 10))
	conn.SetConnMaxLifetime(time.Duration(getEnvInt("DB_CONN_MAX_LIFETIME", 180000)))

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

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")
	return dbConn, err
}

func getStrDB() string {
	dbHost := os.Getenv("DB_HOST")
	dbPort := getEnvInt("DB_PORT", port)
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	strDb := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	return strDb
}

func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	i, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return i
}
