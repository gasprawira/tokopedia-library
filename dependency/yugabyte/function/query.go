package yugabyte

import (
	"github.com/yugabyte/gocql"

	yb "github.com/gasprawira/tokopedia-library/dependency/yugabyte"
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
