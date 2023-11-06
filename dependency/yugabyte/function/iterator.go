package yugabyte

import "github.com/yugabyte/gocql"

type iterator struct {
	i *gocql.Iter
}

func (i *iterator) MapScan(result map[string]interface{}) bool {
	return i.i.MapScan(result)
}

func (i *iterator) Close() error {
	return i.i.Close()
}
