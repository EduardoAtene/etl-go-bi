// internal/infrastructure/database/database.go

package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/EduardoAtene/etl-go-bi/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLConnection struct {
	Conn *sql.DB
}

func NewMySQLConnection(cfg config.DatabaseConfig) (*MySQLConnection, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to MySQL database.")
	return &MySQLConnection{Conn: db}, nil
}

func (db *MySQLConnection) Begin() (*sql.Tx, error) {
	return db.Conn.Begin()
}

func (db *MySQLConnection) Close() error {
	return db.Conn.Close()
}
