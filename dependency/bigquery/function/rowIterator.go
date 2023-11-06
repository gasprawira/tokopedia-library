package bigquery

import "cloud.google.com/go/bigquery"

type rowIterator struct {
	r *bigquery.RowIterator
}

func (r rowIterator) Next(dst interface{}) error {
	return r.r.Next(dst)
}
