package bq

import "context"

// Package bq contains interfaces wrapped from bigquery library.
//
// The purpose of this package is to reach 100% coverage on bigquery repositories
// Add another function if needed

// Client may be used to perform BigQuery operations.
type Client interface {
	// Query creates a query with string q.
	// The returned Query may optionally be further configured before its Run method is called.
	Query(q string) Query

	// IsClientNil check wether the bigquery client is nil or not
	IsClientNil() bool
}

// Query queries data from a BigQuery table. Use Client.Query to create a Query.
type Query interface {
	// Read submits a query for execution and returns the results via a RowIterator.
	// It is a shorthand for Query.Run followed by Job.Read.
	Read(ctx context.Context) (RowIterator, error)
}

// RowIterator provides access to the result of a BigQuery lookup.
type RowIterator interface {
	// Next loads the next row into dst. Its return value is iterator.Done if there
	// are no more results. Once Next returns iterator.Done, all subsequent calls
	// will return iterator.Done.
	Next(dst interface{}) error
}
