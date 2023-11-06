package yugabyte

import (
	"github.com/gasprawira/tokopedia-library/dependency/yb"

	"github.com/yugabyte/gocql"
)

type client struct {
	s *gocql.Session
}

func New(s *gocql.Session) yb.Session {
	return &client{
		s: s,
	}
}

// IsClientNil implements yb.Client
func (c *client) IsClientNil() bool {
	return c.s == nil
}

// Query implements yb.Client
func (c *client) Query(statement string, arguments []interface{}) yb.Query {
	return &query{
		q: c.s.Query(statement, arguments...),
	}
}

// Batch implements yb.NewBatch
func (c *client) Batch(batchType yb.BatchType) yb.Batch {
	return &batch{
		b: c.s.NewBatch(gocql.BatchType(batchType)),
		s: c.s,
	}
}
