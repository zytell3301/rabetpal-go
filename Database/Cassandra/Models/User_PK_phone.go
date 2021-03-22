package Models

import (
	"github.com/gocql/gocql"
	"rabetpal/Database/Cassandra"
)

var User_PK_phoneMetaData = Cassandra.TableMetaData{
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
	Pk:       map[string]struct{}{"phone": {}},
	Ck:       nil,
	Keyspace: "rabetpal",
}

func NewUserPKPhone(values map[string]interface{}, statement *gocql.Batch) bool {
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
