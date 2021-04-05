package mysql

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDB(conn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", conn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s: %w", conn, err)
	}
	db.SetConnMaxLifetime(30 * time.Second)
	db.SetMaxOpenConns(150)
	db.SetMaxIdleConns(15)

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
