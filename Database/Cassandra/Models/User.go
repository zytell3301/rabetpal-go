package Models

import (
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
	switch statement == nil {
	case true:
		statement = Cassandra.ConnectionManager.GetSession("rabetpal").NewBatch(gocql.LoggedBatch)
	}
	Cassandra.AddId(&values)
	Cassandra.NewRecord("users", values, statement, UsersMetaData)
	Cassandra.AddDependencies(UsersMetaData.DependsOn, values, statement)
	error := Cassandra.ConnectionManager.GetSession("rabetpal").ExecuteBatch(statement)
	switch error != nil {
	case true:
		return false
	}
	return true
}
