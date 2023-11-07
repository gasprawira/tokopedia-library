package yugabyte

import (
	"time"

	"github.com/yugabyte/gocql"
)

const (
	InsertData     = "insert"
	UpdateData     = "update"
	AppendData     = "append"
	PrependData    = "prepend"
	RemoveData     = "remove"
	SelectData     = "select"
	DeleteData     = "delete"
	CreateKeyspace = "create keyspace"
	CreateTable    = "create table"
	TruncateTable  = "truncate table"
	DropTable      = "drop table"
	Consistency    = gocql.Quorum
	Timeout        = time.Second * 100
	ConnectTimeout = time.Second * 100
)

// ConsistencyMap for mapping consistency to gocql.Consistency
var (
	ConsistencyMap = map[string]gocql.Consistency{
		"Any":         gocql.Any,
		"One":         gocql.One,
		"Two":         gocql.Two,
		"Three":       gocql.Three,
		"Quorum":      gocql.Quorum,
		"All":         gocql.All,
		"LocalQuorum": gocql.LocalQuorum,
		"EachQuorum":  gocql.EachQuorum,
		"LocalOne":    gocql.LocalOne,
	}

	EqualGoDataType = map[string]string{
		"int":      "int",
		"[]int":    "list<int>",
		"int64":    "bigint",
		"[]int64":  "list<bigint>",
		"string":   "text",
		"[]string": "list<text>",
	}
)
