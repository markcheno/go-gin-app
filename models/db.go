package models

import (
	"github.com/jinzhu/gorm"
	// postgress db driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// sqlite db driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB abstraction
type DB struct {
	*gorm.DB
}

// NewPostgresDB - postgres database
func NewPostgresDB(dataSourceName string) (*DB, error) {

	db, err := gorm.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.DB().Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

// NewSqliteDB - sqlite database
func NewSqliteDB(databaseName string) (*DB, error) {

	db, err := gorm.Open("sqlite3", databaseName)
	if err != nil {
		return nil, err
	}

	if err = db.DB().Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
