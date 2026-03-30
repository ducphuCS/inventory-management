package storage

import (
	"database/sql"
	"fmt"

	"github.com/ducphu/inventory-service/models"
	_ "modernc.org/sqlite"
)

// SQLiteStorage is the data access layer for the inventory-service.
// It manages the connection to an SQLite database.
type SQLiteStorage struct {
	db *sql.DB
}

// NewSQLiteStorage initializes a new connection to an SQLite database.
// It returns a pointer to the storage handler and ensures the schema is ready.
func NewSQLiteStorage(dbPath string) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	dbStorage := &SQLiteStorage{db: db}
	if err := dbStorage.initSchema(); err != nil {
		return nil, fmt.Errorf("failed to initialize schema: %w", err)
	}

	return dbStorage, nil
}

// initSchema creates the items table if it doesn't already exist.
// This ensures the database is correctly structured for operations.
func (dbStorage *SQLiteStorage) initSchema() error {
	query := `
	CREATE TABLE IF NOT EXISTS items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		product_name TEXT NOT NULL,
		stock_count INTEGER NOT NULL,
		item_details TEXT
	);`
	_, err := dbStorage.db.Exec(query)
	return err
}

// AddItem inserts a new product entry into the inventory table.
// It returns the newly created item's unique identifier.
func (dbStorage *SQLiteStorage) AddItem(item models.Item) (int, error) {
	res, err := dbStorage.db.Exec("INSERT INTO items (product_name, stock_count, item_details) VALUES (?, ?, ?)", item.ProductName, item.StockCount, item.ItemDetails)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// GetTotalCount returns the total number of product records in the inventory.
func (dbStorage *SQLiteStorage) GetTotalCount() (int, error) {
	var count int
	err := dbStorage.db.QueryRow("SELECT COUNT(*) FROM items").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetItemByID retrieves a single product record from the inventory by its ID.
func (dbStorage *SQLiteStorage) GetItemByID(id int) (*models.Item, error) {
	var item models.Item
	query := "SELECT id, product_name, stock_count, item_details FROM items WHERE id = ?"
	err := dbStorage.db.QueryRow(query, id).Scan(&item.ID, &item.ProductName, &item.StockCount, &item.ItemDetails)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Not found
		}
		return nil, err
	}
	return &item, nil
}

// Close closes the connection to the SQLite database.
func (dbStorage *SQLiteStorage) Close() error {
	return dbStorage.db.Close()
}
