package yb

// Package yb contains interfaces wrapped from yugabyte gocql library.
//
// The purpose of this package is to reach 100% coverage on yugabyte repositories
// Add another function if needed

//go:generate mockgen -source=interface.go -destination=mock/mock.go -package=mock

// Session - is used to wrap yugabyte gocql Session
type Session interface {
	// Query creates query with string statement and array of interface argument
	Query(statement string, arguments []interface{}) Query

	// Batch creates batch with batch type
	Batch(batchType BatchType) Batch

	// IsClientNil check wether the yugabyte client is nil or not
	IsClientNil() bool
}

// Query queries data from a yugabyte table. Use Client.Query to create a Query.
type Query interface {
	// Iter executes the query and returns an iterator capable of iterating over all results.
	Iter() Iter

	// Exec executes the query without returning any rows.
	Exec() error
}

type Iter interface {
	// MapScan takes a map[string]interface{} and populates it with a row that is returned from cassandra.
	MapScan(map[string]interface{}) bool

	// Close closes the iterator and returns any errors that happened during the query or the iteration.
	Close() error
}

type Batch interface {
	// AddQuery - adds new query to batch
	AddQuery(statement string, arguments ...interface{})

	// ExecuteBatch - executes the given batch
	ExecuteBatch() error
}
