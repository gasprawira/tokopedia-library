package yugabyte

import "github.com/yugabyte/gocql"

// InterfaceYugabyte for interfacing YugabyteMethods
type InterfaceYugabyte interface {
	Insert(data YugabyteData, metricInfo []string) (err error)
	Update(data YugabyteData, metricInfo []string) (err error)
	Prepend(data YugabyteData, metricInfo []string) (err error)
	Append(data YugabyteData, metricInfo []string) (err error)
	Remove(data YugabyteData, metricInfo []string) (err error)
	Get(data YugabyteData, metricInfo []string) (result []map[string]interface{}, err error)
	Delete(data YugabyteData, metricInfo []string) (err error)
	CreateKeyspace(data YugabyteMeta) (err error)
	CreateTable(data YugabyteMeta) (err error)
	TruncateTable(data YugabyteMeta) (err error)
	DropTable(data YugabyteMeta) (err error)
	CloseSession()
	GetSession() (session *gocql.Session)
}
