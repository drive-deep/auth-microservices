package database

import "context"

// Database interface defines the operations for any database (Insert, Update, Delete, Get, and GetAll)
type Database interface {
	Insert(ctx context.Context, value interface{}) error
	Update(ctx context.Context, key string, value interface{}) error
	Delete(ctx context.Context, key string) error
	Get(ctx context.Context, key string, result interface{}) error
	GetAll(ctx context.Context, model interface{}, limit, offset int) ([]interface{}, error)
}
