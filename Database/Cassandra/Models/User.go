package Models

import (
	"fmt"
	"github.com/gocql/gocql"
	"rabetpal/Database/Cassandra"
)

type User struct {
	id            gocql.UUID
	name          string
	lastname      string
	username      string
	mobile        string
	phone         string
	address       map[string]string
	bank_accounts map[string]string
	national_code string
	profile_pic   string
	email         string
	user_type     string
	balance       int
	admin         map[string]interface{}
	seller        map[string]interface{}
	broker        map[string]interface{}
	password      string
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
	DependsOn: Cassandra.TableDependency{newUserPKMobile},
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
