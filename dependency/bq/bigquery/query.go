package bigquery

import (
	"context"

	"cloud.google.com/go/bigquery"

	"github.com/gasprawira/tokopedia-library/dependency/bq"
)

type query struct {
	q *bigquery.Query
}

func (q query) Read(ctx context.Context) (iter bq.RowIterator, err error) {
	r, err := q.q.Read(ctx)
	iter = rowIterator{r: r}
	return iter, err
}
