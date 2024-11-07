package database

import (
	"context"
	"fmt"
	"log"

	"github.com/go-pg/pg/v10"
	// Import the db package to implement the interface
)

// PostgresDatabase struct represents a connection to a PostgreSQL database
type PostgresDatabase struct {
	db *pg.DB
}

func NewPostgresDatabase(addr, user, password, dbName string) (*PostgresDatabase, error) {
	db := pg.Connect(&pg.Options{
		Addr:     addr,
		User:     user,
		Password: password,
		Database: dbName,
	})

	// Check connection
	if err := db.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return &PostgresDatabase{db: db}, nil
}

// InitPostgres initializes the PostgreSQL connection and ensures the connection is active
func InitPostgres(addr, user, password, dbName string) (*PostgresDatabase, error) {
	// Initialize the PostgreSQL database connection
	db, err := NewPostgresDatabase(addr, user, password, dbName)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize PostgreSQL: %v", err)
	}

	// Ping the database to verify connection
	_, errDb := db.db.QueryOne(pg.Scan(&struct{}{}), "SELECT 1")
	if errDb != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", errDb)
	}

	log.Println("Postgres initialized and successfully connected")

	// Return the PostgresDatabase instance
	return db, nil
}

func (p *PostgresDatabase) Close() error {
	return p.db.Close()
}

// Insert implements the Insert method of the Database interface
func (p *PostgresDatabase) Insert(ctx context.Context, value interface{}) error {
	_, err := p.db.Model(value).Insert()
	if err != nil {
		return fmt.Errorf("failed to insert data: %v", err)
	}
	return nil
}

// Update implements the Update method of the Database interface
func (p *PostgresDatabase) Update(ctx context.Context, key string, value interface{}) error {
	_, err := p.db.Model(value).Where("key = ?", key).Update()
	if err != nil {
		return fmt.Errorf("failed to update data: %v", err)
	}
	return nil
}

// Delete implements the Delete method of the Database interface
// Delete implements the Delete method of the Database interface
func (p *PostgresDatabase) Delete(ctx context.Context, key string, model interface{}) error {
	// Delete the record where the key matches
	_, err := p.db.Model(model).Where("key = ?", key).Delete()
	if err != nil {
		return fmt.Errorf("failed to delete data: %v", err)
	}
	return nil
}

// Get implements the Get method of the Database interface
func (p *PostgresDatabase) Get(ctx context.Context, key string, result interface{}) error {
	err := p.db.Model(result).Where("key = ?", key).Select()
	if err != nil {
		return fmt.Errorf("failed to retrieve data: %v", err)
	}
	return nil
}

// GetAll retrieves all records of the given model with optional pagination
func (p *PostgresDatabase) GetAll(ctx context.Context, model interface{}, limit, offset int) ([]interface{}, error) {
	var results []interface{}

	query := p.db.Model(model)

	// Apply pagination if limit is provided
	if limit > 0 {
		query = query.Limit(limit)
	}

	// Apply offset if it's provided
	if offset > 0 {
		query = query.Offset(offset)
	}

	err := query.Select()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve all data: %v", err)
	}

	return results, nil
}

// CreateSchema creates the database schema by generating tables for all models passed
// func (p *PostgresDatabase) CreateSchema(models []interface{}) error {
// 	for _, model := range models {
// 		err := p.db.Model(model).CreateTable(&orm.CreateTableOptions{
// 			IfNotExists: true, // Prevent errors if the table already exists
// 		})
// 		if err != nil {
// 			return fmt.Errorf("failed to create schema for model %v: %v", model, err)
// 		}
// 	}

// 	log.Println("Database schema created successfully")
// 	return nil
// }

// // MigrateSchema runs migrations on the database schema
// // This function is just an example of how you could run schema changes
// func (p *PostgresDatabase) MigrateSchema(models []interface{}) error {
// 	// For simplicity, this will attempt to create the table, adding the "IfNotExists" clause.
// 	// You could expand this by implementing more complex migration logic.
// 	for _, model := range models {
// 		err := p.db.Model(model).CreateTable(&orm.CreateTableOptions{
// 			IfNotExists: false, // Assume we want to apply changes (not "IfNotExists")
// 		})
// 		if err != nil {
// 			return fmt.Errorf("failed to migrate schema for model %v: %v", model, err)
// 		}
// 	}

// 	log.Println("Database schema migrated successfully")
// 	return nil
// }

// // DropSchema drops the tables for all models passed
// func (p *PostgresDatabase) DropSchema(models []interface{}) error {
// 	for _, model := range models {
// 		_, err := p.db.Model(model).DropTable(&orm.DropTableOptions{
// 			IfExists: true, // Prevent errors if the table doesn't exist
// 		})
// 		if err != nil {
// 			return fmt.Errorf("failed to drop schema for model %v: %v", model, err)
// 		}
// 	}

// 	log.Println("Database schema dropped successfully")
// 	return nil
// }
