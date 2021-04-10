package Models

import (
	"github.com/gocql/gocql"
	"rabetpal/Database/Cassandra"
)

var UserPKPhoneMetaData = Cassandra.TableMetaData{
	Table: "users_pk_mobile",
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
	statement := Cassandra.ConnectionManager.GetSession(UserPKPhoneMetaData.Keyspace).NewBatch(gocql.LoggedBatch)
	Cassandra.AddId(&values)
	newUserPKPhone(values, statement)
	error := Cassandra.ConnectionManager.GetSession(UserPKPhoneMetaData.Keyspace).ExecuteBatch(statement)
	switch error != nil {
	case true:
		return false
	}
	return true
}

func newUserPKPhone(values map[string]interface{}, statement *gocql.Batch) bool {
	switch Cassandra.CheckPK(UserPKPhoneMetaData,&values) {
	case false:
		return false
	}
	return Cassandra.NewRecord(UserPKPhoneMetaData.Table, values, statement, UserPKPhoneMetaData)
}
