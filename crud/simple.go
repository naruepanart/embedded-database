package crud

import (
	"fmt"
	"github.com/cockroachdb/pebble"
)

// CreateKeyValue adds a new key-value pair to the Pebble database.
func CreateKeyValue(db *pebble.DB, key, value []byte) error {
	// Check if the key or value is empty
	if len(key) == 0 {
		return fmt.Errorf("key must not be empty")
	}
	if len(value) == 0 {
		return fmt.Errorf("value must not be empty")
	}

	// Set the key-value pair in the database
	db.Set(key, value, nil)
	return nil
}

// ReadKeyValue retrieves the value associated with a given key from the Pebble database.
func ReadKeyValue(db *pebble.DB, key []byte) ([]byte, error) {
	// Check if the key is empty
	if len(key) == 0 {
		return nil, fmt.Errorf("key must not be empty")
	}

	// Retrieve the value for the given key
	value, closer, err := db.Get(key)
	if err != nil {
		return nil, fmt.Errorf("failed to get value for key: %w", err)
	}
	defer closer.Close() // Ensure the closer is closed properly

	return value, nil
}

// UpdateKeyValue updates an existing key-value pair with a new value in the Pebble database.
func UpdateKeyValue(db *pebble.DB, key, newValue []byte) error {
	// Check if the key or value is empty
	if len(key) == 0 {
		return fmt.Errorf("key must not be empty")
	}
	if len(newValue) == 0 {
		return fmt.Errorf("value must not be empty")
	}

	// Update the key-value pair in the database
	db.Set(key, newValue, nil)
	return nil
}

// DeleteKeyValue removes a key-value pair associated with a given key from the Pebble database.
func DeleteKeyValue(db *pebble.DB, key []byte) error {
	// Check if the key is empty
	if len(key) == 0 {
		return fmt.Errorf("key must not be empty")
	}

	// Delete the key-value pair from the database
	db.Delete(key, nil)
	return nil
}
