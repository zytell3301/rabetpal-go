package Keyspaces

import (
	"fmt"
	"github.com/gocql/gocql"
	"rabetpal/Database/Cassandra"
	"time"
)

var connection = Cassandra.Connection{

}

func init() {
	fmt.Println("The connection is being set")
	connection.Cluster = gocql.NewCluster("127.0.0.1")
	connection.Cluster.Consistency = gocql.One
	connection.Cluster.Timeout = 1 * time.Second
	connection.Cluster.Keyspace = "rabetpal"
	connection.Session, _ = connection.Cluster.CreateSession()
	Cassandra.ConnectionManager.AddSession("rabetpal", connection)
}
