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
	Table     string
	Columns   map[string]struct{}
	Pk        map[string]struct{}
	Ck        map[string]struct{}
	Keyspace  string
	DependsOn TableDependency
	Maps map[string]interface{}
}

type TableDependency []func(map[string]interface{}, *gocql.Batch) bool

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

func FilterData(data map[string]interface{}, metaData TableMetaData) (map[string]interface{}) {
	values := make(map[string]interface{})
	for column, _ := range metaData.Columns {
		value, isset := data[column]
		switch isset {
		case true:
			values[column] = value
		}
	}
	return values
}

func GenerateEmptyInputs(count int) string {
	var inputs []string
	for i := 0; i < count; i++ {
		inputs = append(inputs, "?")
	}
	return strings.Join(inputs, ",")
}

func BindArgs(data map[string]interface{}) ([]interface{},string) {
	Args := []interface{}{}
	fields := make([]string,0)
	for field, value := range data {
		Args = append(Args, value)
		fields = append(fields,field)
	}
	return Args,strings.Join(fields,",")
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
	data := FilterData(values, metaData)
	switch len(data) == 0 {
	case true:
		return false
	}
	Args,fields := BindArgs(data)
	batch.Entries = append(batch.Entries, gocql.BatchEntry{
		Stmt:       "INSERT INTO " + table + " (" + fields + ") VALUES (" + GenerateEmptyInputs(len(data)) + ")",
		Args:       Args,
		Idempotent: false,
	})
	return true
}

func AddDependencies(dependencies TableDependency, values map[string]interface{}, statement *gocql.Batch) bool {
	isSuccessful := true
	for _, value := range dependencies {
		isSuccessful = isSuccessful && value(values, statement)
	}
	return isSuccessful
}

func CheckPK(metaData TableMetaData, data *map[string]interface{}) bool {
	for field := range metaData.Pk {
		switch _, isSet := (*data)[field]; isSet {
		case false:
			return false
		}
	}
	return true
}
