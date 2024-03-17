package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

var cfg *config

type config struct {
	DB DBConfig
}

type ApiConfig struct {
	Port string
}

type DBConfig struct {
	Database string
}

func init() {
	cfg = &config{
		DB: DBConfig{
			Database: os.Getenv("DATABASE"),
		},
	}
}

func Load() error {
	cfg = new(config)
	cfg.DB = DBConfig{
		Database: os.Getenv("DATABASE"),
	}
	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func NewConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		fmt.Println("Erro ao abrir o banco de dados:", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("Erro ao fechar o banco de dados:", err)
		}
	}(db)

	// Verificar se a conexão está disponível
	err = db.Ping()
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
	}

	fmt.Println("Conexão com o banco de dados SQLite estabelecida com sucesso!")
	return db, nil
}
