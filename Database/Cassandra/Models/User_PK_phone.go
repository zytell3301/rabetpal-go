package Models

import (
	"github.com/gocql/gocql"
	"rabetpal/Database/Cassandra"
)

var UserPKPhoneMetaData = Cassandra.TableMetaData{
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

func NewUserPKPhone(values map[string]interface{}) bool {
	statement := Cassandra.ConnectionManager.GetSession("rabetpal").NewBatch(gocql.LoggedBatch)
	Cassandra.AddId(&values)
	newUserPKPhone(values, statement)
	error := Cassandra.ConnectionManager.GetSession("rabetpal").ExecuteBatch(statement)
	switch error != nil {
	case true:
		return false
	}
	return true
}

func newUserPKPhone(values map[string]interface{}, statement *gocql.Batch) bool {
	return Cassandra.NewRecord("users_pk_mobile", values, statement, UserPKPhoneMetaData)
}
