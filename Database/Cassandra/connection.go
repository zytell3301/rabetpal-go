package Cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	"strings"
)

type Connection struct {
	Cluster *gocql.ClusterConfig
	Session *gocql.Session
}

type TableMetaData struct {
	Columns   map[string]struct{}
	Pk        map[string]struct{}
	Ck        map[string]struct{}
	Keyspace  string
	DependsOn []func(map[string]interface{},*gocql.Batch) bool
}

var connections = make(map[string]Connection)

var ConnectionManager = struct {
	GetSession func(name string) *gocql.Session
	AddSession func(keyspace string, connection Connection)
}{
	GetSession: func(name string) *gocql.Session {
		return connections[name].Session
	},
	AddSession: func(keyspace string, connection Connection) {
		fmt.Println("Connection " + keyspace + " just got set")
		connections[keyspace] = connection
	},
}

func FilterData(data map[string]interface{}, metaData TableMetaData) (map[string]interface{}, []string) {
	var fields []string
	values := make(map[string]interface{})
	for column, _ := range metaData.Columns {
		value, isset := data[column]
		switch isset {
		case true:
			values[column] = value
			fields = append(fields, column)
		}
	}
	return values, fields
}

func GenerateEmptyInputs(count int) string {
	var inputs []string
	for i := 0; i < count; i++ {
		inputs = append(inputs, "?")
	}
	return strings.Join(inputs, ",")
}

func BindArgs(data map[string]interface{}) []interface{} {
	Args := []interface{}{}
	for _, value := range data {
		Args = append(Args, value)
	}
	return Args
}

func AddId(values *map[string]interface{}) {
	id, _ := gocql.RandomUUID()
	_, isset := (*values)["id"]
	switch isset {
	case false:
		(*values)["id"] = id.String()
	}
}

func NewRecord(table string, values map[string]interface{}, batch *gocql.Batch, metaData TableMetaData) bool {
	data, fields := FilterData(values, metaData)
	switch len(fields) == 0 {
	case true:
		return false
	}
	batch.Entries = append(batch.Entries, gocql.BatchEntry{
		Stmt:       "INSERT INTO " + table + " (" + strings.Join(fields, ",") + ") VALUES (" + GenerateEmptyInputs(len(fields)) + ")",
		Args:       BindArgs(data),
		Idempotent: false,
	})
	return true
}
