package yugabyte

import (
	"time"

	"github.com/yugabyte/gocql"
)

type (
	YugabyteInstance struct {
		Session  *gocql.Session
		Keyspace string
	}

	// YugabyteConfig for configuration yugabyte session
	YugabyteConfig struct {
		ClusterDSN               []string
		Keyspace                 string
		Environment              string
		Port                     int
		Consistency              string
		Timeout                  time.Duration
		ConnectTimeout           time.Duration
		DisableInitialHostLookup bool
		YugabyteUsername         string
		YugabytePassword         string
		Keepalive                time.Duration
		LocalDC                  string
		Version                  string
		MaxConnection            int
	}

	// YugabyteData for yugabyte data DML
	YugabyteData struct {
		TableName     string
		PrimaryKey    map[string]interface{}
		ClusteringKey map[string]interface{}
		// DataFields is the fields for inserting or update to Yugabyte
		DataFields map[string]interface{}
		// FilteredFields is the fields that you want to filter (Where Clause)
		FilteredFields map[string]interface{}
		// Operator is math operator to filter the fields in Where Clause
		// Only for symbol ('<', '<=', '>', '>=', '<>') because the default is using '='
		Operator map[string]string
		// ColumnFields is the fields that you want to Select or Remove
		ColumnFields     []string
		ConsistencyLevel string
		TimeToLive       int
		// Limit is the fields used by Get/Select to limit the output
		Limit int
		// MapReferenceKey is the field used to filter map reference key
		// The value is keys on the maps, that equivalent with DataFields
		// ex: column_name = fruits, value = {'orange': 1, 'melon': 3}, so, in MapReferenceKey = {'fruits': []string{'orange', 'melon'}} and DataFields = {'orange': 1, 'melon': 3}
		MapReferenceKey map[string][]string
	}

	// YugabyteMeta for yugabyte DDL
	YugabyteMeta struct {
		ReplicationStrategy string
		CompactionStrategy  string
		TableName           string
		DataFieldsType      map[string]string
		QueryStatement      string
		ConsistencyLevel    string
		PrimaryKeys         []string
		ClusteringKeysOrder map[string]string
	}
)
