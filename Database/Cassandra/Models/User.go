package Models

import (
	"fmt"
	"github.com/gocql/gocql"
	"rabetpal/Database/Cassandra"
)

type User struct {
	id       gocql.UUID
	name     string
	lastname string
}

var UsersMetaData = Cassandra.TableMetaData{
	Table: "users",
	Columns: map[string]struct{}{
		"id":            {},
		"name":          {},
		"lastname":      {},
		"username":      {},
		"mobile":        {},
		"phone":         {},
		"address":       {},
		"bank_accounts": {},
		"national_code": {},
		"profile_pic":   {},
		"email":         {},
		"user_type":     {},
		"balance":       {},
		"admin":         {},
		"seller":        {},
		"broker":        {},
		"password":      {},
	},
	Pk:        map[string]struct{}{"id": {}},
	Ck:        nil,
	Keyspace:  "rabetpal",
	DependsOn: Cassandra.TableDependency{newUserPKPhone},
}

func NewUser(values map[string]interface{}, statement *gocql.Batch) bool {
	connection := Cassandra.ConnectionManager.GetSession(UsersMetaData.Keyspace)
	switch statement == nil {
	case true:
		statement = connection.NewBatch(gocql.LoggedBatch)
	}
	Cassandra.AddId(&values)
	Cassandra.NewRecord(UsersMetaData.Table, values, statement, UsersMetaData)
	Cassandra.AddDependencies(UsersMetaData.DependsOn, values, statement)
	error := connection.ExecuteBatch(statement)
	switch error != nil {
	case true:
		fmt.Println(error)
		return false
	}
	return true
}