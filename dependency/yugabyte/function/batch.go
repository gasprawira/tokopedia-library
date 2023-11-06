package yugabyte

import "github.com/yugabyte/gocql"

type batch struct {
	b *gocql.Batch
	s *gocql.Session
}

// AddQuery - adds new query to batch
func (b *batch) AddQuery(statement string, arguments ...interface{}) {
	b.b.Query(statement, arguments...)
}

// ExecuteBatch - executes the given batch
func (b *batch) ExecuteBatch() error {
	return b.s.ExecuteBatch(b.b)
}
