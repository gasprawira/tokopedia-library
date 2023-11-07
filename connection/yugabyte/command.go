package yugabyte

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/yugabyte/gocql"
)

// Insert for inserting data
func (i *YugabyteInstance) Insert(data YugabyteData, metricInfo []string) (err error) {
	_, err = i.ExecuteDML(data, InsertData, metricInfo)
	return
}

// Update for updating data
func (i *YugabyteInstance) Update(data YugabyteData, metricInfo []string) (err error) {
	_, err = i.ExecuteDML(data, UpdateData, metricInfo)
	return
}

// Prepend for adding data in front
func (i *YugabyteInstance) Prepend(data YugabyteData, metricInfo []string) (err error) {
	_, err = i.ExecuteDML(data, PrependData, metricInfo)
	return
}

// Append for adding data behind
func (i *YugabyteInstance) Append(data YugabyteData, metricInfo []string) (err error) {
	_, err = i.ExecuteDML(data, AppendData, metricInfo)
	return
}

// Remove for remove data in list
func (i *YugabyteInstance) Remove(data YugabyteData, metricInfo []string) (err error) {
	_, err = i.ExecuteDML(data, RemoveData, metricInfo)
	return
}

// Get Data from Yugabyte
func (i *YugabyteInstance) Get(data YugabyteData, metricInfo []string) (result []map[string]interface{}, err error) {
	result, err = i.ExecuteDML(data, SelectData, metricInfo)
	return
}

// Delete Data from Yugabyte
func (i *YugabyteInstance) Delete(data YugabyteData, metricInfo []string) (err error) {
	_, err = i.ExecuteDML(data, DeleteData, metricInfo)
	return
}

// CreateKeyspace for creating a keyspace
func (i *YugabyteInstance) CreateKeyspace(data YugabyteMeta) (err error) {
	return i.ExecuteDDL(data, CreateKeyspace)
}

// CreateTable for creating a table
func (i *YugabyteInstance) CreateTable(data YugabyteMeta) (err error) {
	return i.ExecuteDDL(data, CreateTable)
}

// TruncateTable for truncating a table
func (i *YugabyteInstance) TruncateTable(data YugabyteMeta) (err error) {
	return i.ExecuteDDL(data, TruncateTable)
}

// DropTable for dropping a table
func (i *YugabyteInstance) DropTable(data YugabyteMeta) (err error) {
	return i.ExecuteDDL(data, DropTable)
}

func (i *YugabyteInstance) CloseSession() {
	i.Session.Close()
}

func (i *YugabyteInstance) GetSession() (session *gocql.Session) {
	return i.Session
}

func (i *YugabyteInstance) ExecuteDML(data YugabyteData, dml string, metricInfo []string) (result []map[string]interface{}, err error) {
	status := "success"

	defer func() {
		if err != nil {
			status = "failed"
		}
		tags := []string{"keyspace:" + i.Keyspace, "table:" + data.TableName, "command:" + dml, "status:" + status}
		tags = append(tags, metricInfo...)
		log.Println("send metric", tags)
	}()

	err = validateDML(data, dml)
	if err != nil {
		return nil, err
	}

	queryStmt, queryArgs := i.generateQueryAndArgs(data, dml)
	query := i.Session.Query(queryStmt, queryArgs...)
	if data.ConsistencyLevel != "" {
		query = query.Consistency(ConsistencyMap[data.ConsistencyLevel])
	}

	if dml == SelectData {
		iterator := query.Iter()
		tempData := make(map[string]interface{})
		for iterator.MapScan(tempData) {
			result = append(result, tempData)
			tempData = make(map[string]interface{})
		}

		err = iterator.Close()
	} else {
		err = query.Exec()
	}
	return
}

func (i *YugabyteInstance) ExecuteDDL(data YugabyteMeta, ddl string) (err error) {
	err = validateDDL(data, ddl)
	if err != nil {
		return
	}

	queryStmt := i.generateQueryDDL(data, ddl)
	query := i.Session.Query(queryStmt)
	if (ddl == TruncateTable || ddl == DropTable) && data.ConsistencyLevel != "" {
		query = query.Consistency(ConsistencyMap[data.ConsistencyLevel])
	}

	err = query.Exec()
	return
}

func (i *YugabyteInstance) generateQueryAndArgs(data YugabyteData, crud string) (queryStmt string, args []interface{}) {
	var columnName []string
	var primaryKey []string
	var clusteringKey []string
	var qm []string

	var ttl string
	if data.TimeToLive != 0 {
		ttl = ` USING TTL ` + strconv.Itoa(data.TimeToLive)
	}

	keys := make([]string, 0)
	for k := range data.DataFields {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, name := range keys {
		tempColumnName := ""
		switch crud {
		case InsertData:
			tempColumnName = name
			qm = append(qm, "?")
		case UpdateData:
			if len(data.MapReferenceKey) > 0 {
				for k, v := range data.MapReferenceKey {
					for _, val := range v {
						if name == val {
							tempColumnName = k + `['` + val + `'] = ?`
						}
					}
				}
			} else {
				tempColumnName = name + " = ?"
			}
		case PrependData:
			tempColumnName = name + " = ? + " + name
		case AppendData:
			tempColumnName = name + " = " + name + " + ?"
		case RemoveData:
			tempColumnName = name + " = " + name + " - ?"
		}
		columnName = append(columnName, tempColumnName)
		args = append(args, data.DataFields[name])
	}

	if crud == InsertData {
		queryStmt = `INSERT INTO ` + i.Keyspace + `.` + data.TableName + ` (` + strings.Join(columnName, ",") + `)` + ` VALUES (` + strings.Join(qm, ",") + `)` + ttl + `;`
		return queryStmt, args
	}

	for pk, value := range data.PrimaryKey {
		concat := pk + " = ?"
		primaryKey = append(primaryKey, concat)
		args = append(args, value)
	}

	for ck, value := range data.ClusteringKey {
		concat := ck + " = ?"
		clusteringKey = append(clusteringKey, concat)
		args = append(args, value)
	}

	switch crud {
	case UpdateData, PrependData, AppendData, RemoveData:
		queryStmt = `UPDATE ` + i.Keyspace + `.` + data.TableName + ttl + ` SET ` + strings.Join(columnName, ", ") + ` WHERE ` + strings.Join(primaryKey, " AND ")
	case SelectData:
		queryStmt = `SELECT ` + strings.Join(data.ColumnFields, ", ") + ` FROM ` + i.Keyspace + `.` + data.TableName + ` WHERE ` + strings.Join(primaryKey, " AND ")
	case DeleteData:
		queryStmt = `DELETE ` + strings.Join(data.ColumnFields, ", ") + ` FROM ` + i.Keyspace + `.` + data.TableName + ` WHERE ` + strings.Join(primaryKey, " AND ")
	}

	if len(data.ClusteringKey) > 0 {
		queryStmt += ` AND ` + strings.Join(clusteringKey, ", ")
	}

	if len(data.FilteredFields) > 0 {
		fields, arg := filteredFields(data)
		queryStmt += ` AND ` + strings.Join(fields, " AND ")
		args = append(args, arg...)
	}

	if data.Limit > 0 {
		queryStmt += ` LIMIT ` + strconv.Itoa(data.Limit)
	}

	queryStmt += `;`
	return queryStmt, args
}

func (i *YugabyteInstance) generateQueryDDL(data YugabyteMeta, ddl string) (queryStmt string) {
	switch ddl {
	case CreateKeyspace:
		return `CREATE KEYSPACE IF NOT EXISTS ` + i.Keyspace + ` WITH REPLICATION = ` + data.ReplicationStrategy + `;`
	case TruncateTable:
		return `TRUNCATE ` + i.Keyspace + `.` + data.TableName + `;`
	case DropTable:
		return `DROP TABLE ` + i.Keyspace + `.` + data.TableName + `;`
	}

	var strategies []string
	var dataFieldsType []string
	var clusteringKeysOrder []string

	for field, goDataType := range data.DataFieldsType {
		if dbDataType, ok := EqualGoDataType[goDataType]; ok {
			dataFieldsType = append(dataFieldsType, field+" "+dbDataType)
		} else {
			dataFieldsType = append(dataFieldsType, field+" "+goDataType)
		}
	}

	queryStmt = `CREATE TABLE IF NOT EXISTS ` + i.Keyspace + `.` + data.TableName + ` (` + strings.Join(dataFieldsType, ",") + `, PRIMARY KEY (` + strings.Join(data.PrimaryKeys, ",") + `)` + `)`

	for field, order := range data.ClusteringKeysOrder {
		clusteringKeysOrder = append(clusteringKeysOrder, field+" "+order)
	}

	if len(clusteringKeysOrder) > 0 {
		strategies = append(strategies, `CLUSTERING ORDER BY (`+strings.Join(clusteringKeysOrder, ",")+`)`)
	}

	if data.CompactionStrategy != "" {
		strategies = append(strategies, `COMPACTION = `+data.CompactionStrategy)
	}

	if len(strategies) > 0 {
		queryStmt += `WITH ` + strings.Join(strategies, " AND ")
	}
	queryStmt += `;`
	return
}

func validateDML(data YugabyteData, dml string) (err error) {
	if data.TableName == "" {
		err = fmt.Errorf("table name is empty")
		return
	}

	if len(data.PrimaryKey) == 0 && dml != InsertData {
		err = fmt.Errorf("the primary key was not added")
		return
	}

	return
}

func validateDDL(data YugabyteMeta, ddl string) (err error) {
	if data.ReplicationStrategy == "" && ddl == CreateKeyspace {
		err = fmt.Errorf("the replication strategy is empty")
		return
	}

	if data.TableName == "" && (ddl == CreateTable || ddl == TruncateTable || ddl == DropTable) {
		err = fmt.Errorf("table name is empty")
		return
	}

	if len(data.PrimaryKeys) == 0 && ddl == CreateTable {
		err = fmt.Errorf("the primary key was not added")
		return
	}

	return
}

func filteredFields(data YugabyteData) (fields []string, args []interface{}) {
	var concat string
	for key, value := range data.FilteredFields {
		if operator, ok := data.Operator[key]; ok {
			concat = key + " " + operator + " ?"
		} else {
			concat = key + " = ?"
		}
		fields = append(fields, concat)
		args = append(args, value)
	}

	return
}
