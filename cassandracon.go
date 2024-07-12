package main

/*
import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func main() {
	 Create a new cluster configuration
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "system"
	cluster.Consistency = gocql.Quorum

	 Create a new session
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal("Error creating session: ", err)
	}
	defer session.Close()

	Execute a simple query
	var clusterName string
	if err := session.Query("SELECT cluster_name FROM local").Scan(&clusterName); err != nil {
		log.Fatal("Error executing query: ", err)
	}

	fmt.Printf("Cluster Name: %s\n", clusterName)
}
*/
