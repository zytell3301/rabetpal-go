package Models

import (
	"fmt"
	"github.com/gocql/gocql"
	"rabetpal/Database/Cassandra"
	"strings"
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
	Pk:       map[string]struct{}{"id": {}},
	Ck:       nil,
	Keyspace: "rabetpal",
}

func NewUser(values map[string]interface{}) bool {
	statement := Cassandra.ConnectionManager.GetSession("rabetpal").NewBatch(gocql.LoggedBatch)
	Cassandra.AddId(&values)
	fmt.Println(newUser(values, statement))
	error := Cassandra.ConnectionManager.GetSession("rabetpal").ExecuteBatch(statement)
	switch error != nil {
	case true:
		fmt.Println("Query error is: " + error.Error())
	}
	return true
}

func newUser(values map[string]interface{}, batch *gocql.Batch) bool {
	data, fields := Cassandra.FilterData(values, UsersMetaData)
	switch len(fields) == 0 {
	case true:
		return false
	}
	batch.Entries = append(batch.Entries, gocql.BatchEntry{
		Stmt:       "INSERT INTO users (" + strings.Join(fields, ",") + ") VALUES (" + Cassandra.GenerateEmptyInputs(len(fields)) + ")",
		Args:       Cassandra.BindArgs(data),
		Idempotent: false,
	})
	return true
}
