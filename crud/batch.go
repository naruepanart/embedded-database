package crud

import (
	"fmt"
	"github.com/cockroachdb/pebble"
)

// adds a new key-value pair to the database using a batch operation for improved performance.
func BatchCreateKeyValue(db *pebble.DB, key, value []byte) error {
	batch := db.NewBatch()
	defer batch.Close()

	batch.Set(key, value, pebble.NoSync)
	batch.Commit(pebble.Sync)
	return nil
}

// retrieves the value associated with a key and properly handles the closer.
func BatchReadKeyValue(db *pebble.DB, key []byte) ([]byte, error) {
	value, closer, err := db.Get(key)
	if err != nil {
		return nil, fmt.Errorf("failed to get value for key: %w", err)
	}
	defer closer.Close()
	return value, nil
}

// updates an existing key-value pair with a new value using a batch operation for improved performance.
func BatchUpdateKeyValue(db *pebble.DB, key, newValue []byte) error {
	batch := db.NewBatch()
	defer batch.Close()

	batch.Set(key, newValue, pebble.NoSync)
	batch.Commit(pebble.Sync)
	return nil
}

// removes a key-value pair from the database using a batch operation for improved performance.
func BatchDeleteKeyValue(db *pebble.DB, key []byte) error {
	batch := db.NewBatch()
	defer batch.Close()

	batch.Delete(key, pebble.NoSync)
	batch.Commit(pebble.Sync)
	return nil
}
