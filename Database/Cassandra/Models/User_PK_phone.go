package Models

import (
	"github.com/gocql/gocql"
	"rabetpal/Database/Cassandra"
)

var UserPKMobileMetaData = Cassandra.TableMetaData{
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
	Pk:       map[string]struct{}{"mobile": {}},
	Ck:       nil,
	Keyspace: "rabetpal",
}

func NewUserPKMobile(values map[string]interface{}) bool {
	statement := Cassandra.ConnectionManager.GetSession(UserPKMobileMetaData.Keyspace).NewBatch(gocql.LoggedBatch)
	Cassandra.AddId(&values,nil)
	newUserPKMobile(values, statement)
	error := Cassandra.ConnectionManager.GetSession(UserPKMobileMetaData.Keyspace).ExecuteBatch(statement)
	switch error != nil {
	case true:
		return false
	}
	return true
}

func newUserPKMobile(values map[string]interface{}, statement *gocql.Batch) bool {
	switch Cassandra.CheckPK(UserPKMobileMetaData,&values) {
	case false:
		return false
	}
	return Cassandra.NewRecord(UserPKMobileMetaData.Table, values, statement, UserPKMobileMetaData)
}
