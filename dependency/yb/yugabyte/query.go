package yugabyte

import (
	"github.com/yugabyte/gocql"

	"github.com/gasprawira/tokopedia-library/dependency/yb"
)

type query struct {
	q *gocql.Query
}

func (q *query) Iter() yb.Iter {
	return &iterator{
		i: q.q.Iter(),
	}
}

func (q *query) Exec() error {
	return q.q.Exec()
}
