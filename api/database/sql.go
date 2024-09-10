package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mouzkolit/GOCli/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
)

type DB struct {
	db *bun.DB
}

// Initializes the SQLite database
func InitializeDB() (*DB, error) {
	db, err := NewDB()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}
	if err := db.InitSchema(); err != nil {
		return nil, fmt.Errorf("failed to initialize schema: %w", err)
	}
	return db, nil
}

// Initializes a new database in memory
func NewDB() (*DB, error) {
	sqldb, err := sql.Open("sqlite3", "school.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())
	return &DB{db: db}, nil
}

// Closes the database connection
func (d *DB) Close() error {
	return d.db.Close()
}

// Returns a context with a timeout of 10 seconds
func (d *DB) getContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}

// Creates the base tables for the application
func (d *DB) InitSchema() error {
	ctx := context.Background()
	models := []interface{}{
		(*models.Pupil)(nil),
		(*models.Teacher)(nil),
		(*models.Class)(nil),
		(*models.Course)(nil),
		(*models.Note)(nil),
		(*models.School)(nil),
		(*models.SchoolLogin)(nil),
	}

	for _, model := range models {
		if err := d.createTableIfNotExists(ctx, model); err != nil {
			return fmt.Errorf("failed to init schema: %w", err)
		}
	}
	return nil
}

// Creates a table if it does not exist
func (d *DB) createTableIfNotExists(ctx context.Context, model interface{}) error {
	exists, err := d.db.NewSelect().Model(model).Exists(ctx)
	if err != nil {
		return fmt.Errorf("error checking if table exists for model %T: %w", model, err)
	}

	if !exists {
		_, err := d.db.NewCreateTable().Model(model).IfNotExists().Exec(ctx)
		if err != nil {
			return fmt.Errorf("error creating table for model %T: %w", model, err)
		}
		log.Printf("Table created for model %T", model)
	} else {
		log.Printf("Table already exists for model %T", model)
	}
	return nil
}
