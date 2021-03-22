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
	DependsOn: []func(map[string]interface{}, *gocql.Batch) bool{NewUserPKPhone},
}

func NewUser(values map[string]interface{}, statement *gocql.Batch) bool {
	switch statement == nil {
	case true:
		statement = Cassandra.ConnectionManager.GetSession("rabetpal").NewBatch(gocql.LoggedBatch)
	}
	Cassandra.AddId(&values)
	error := Cassandra.ConnectionManager.GetSession("rabetpal").ExecuteBatch(statement)
	switch error != nil {
	case true:
		return false
	}
	return true
}
