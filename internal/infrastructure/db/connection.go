package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strconv"
)

type DB struct {
	DB *sql.DB
}

var dbConn = &DB{}

func OpenConn() (*DB, error) {

	dbHost := getEnvString("DB_HOST", host)
	dbPort := getEnvInt("DB_PORT", port)
	dbUser := getEnvString("DB_USER", user)
	dbName := getEnvString("DB_NAME", dbname)
	dbPassword := getEnvString("DB_PASSWORD", password)
	strDb := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	conn, err := sql.Open(driverName, strDb)
	conn.SetMaxIdleConns(2)
	conn.SetMaxOpenConns(5)
	conn.SetConnMaxLifetime(180000)

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

func getEnvString(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return i
}
