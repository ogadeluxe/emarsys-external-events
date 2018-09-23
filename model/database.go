package model

import (
	"database/sql"

	"github.com/ogadeluxe/emarsys-external-events/config"
)

// Datastore si an interface
// to define what stuff that a model can do
type Datastore interface {
	AllWishlist() ([]*Wishlist, error)
}

// DB is used to create mysql connection
type DB struct {
	*sql.DB
}

// NewDB : creates and open new connection
func NewDB() (*DB, error) {
	db, err := sql.Open("mysql", config.Items.MySQL)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
