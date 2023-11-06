package bigquery

import (
	"cloud.google.com/go/bigquery"

	bq "github.com/gasprawira/tokopedia-library/dependency/bigquery"
)

type client struct {
	c *bigquery.Client
}

// New create new object for bigquery
func New(c *bigquery.Client) bq.Client {
	return client{c: c}
}

func (c client) Query(q string) bq.Query {
	return query{q: c.c.Query(q)}
}

func (c client) IsClientNil() bool {
	return c.c == nil
}
