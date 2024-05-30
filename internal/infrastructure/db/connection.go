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
	dbHost := getEnvString("DB_HOST", host)
	dbPort := getEnvInt("DB_PORT", port)
	dbUser := getEnvString("DB_USER", user)
	dbName := getEnvString("DB_NAME", dbname)
	dbPassword := getEnvString("DB_PASSWORD", password)

	strDb := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	return strDb
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
