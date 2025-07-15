package models

import (
	"context"
	"errors"
	"time"
)

// DataSyncRequest represents a request to synchronize data between databases
type DataSyncRequest struct {
	SourceDB string `json:"source_db"`
	TargetDB string `json:"target_db"`
	TableName string `json:"table_name"`
}

// DataSyncResponse represents a response after synchronizing data between databases
type DataSyncResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	RowsAffected int `json:"rows_affected"`
}

// CacheItem represents a cached item
type CacheItem struct {
	Key   string `json:"key"`
	Value interface{} `json:"value"`
	TTL   time.Duration `json:"ttl"`
}

// DataSyncError represents an error occurred during data synchronization
type DataSyncError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *DataSyncError) Error() string {
	return e.Message
}

func NewDataSyncError(code int, message string) error {
	return &DataSyncError{Code: code, Message: message}
}

// Validate checks if the DataSyncRequest is valid
func (r *DataSyncRequest) Validate() error {
	if r.SourceDB == "" || r.TargetDB == "" || r.TableName == "" {
		return NewDataSyncError(400, "invalid request: source_db, target_db, and table_name are required")
	}
	return nil
}
