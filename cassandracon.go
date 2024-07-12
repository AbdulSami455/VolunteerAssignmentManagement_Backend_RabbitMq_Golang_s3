package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func cluster_connection() {
	// Connect to Cassandra cluster running on localhost
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "managers_keyspace" // Specify the keyspace to use
	cluster.Consistency = gocql.Quorum

	// Create session for executing queries
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal("Error creating session: ", err)
	}
	defer session.Close()

	// Query to fetch tables in the specified keyspace
	query := "SELECT table_name FROM system_schema.tables WHERE keyspace_name = 'managers_keyspace'"

	iter := session.Query(query).Iter()
	var tableName string

	fmt.Println("Tables in managers_keyspace:")
	for iter.Scan(&tableName) {
		fmt.Println(tableName)
	}
	if err := iter.Close(); err != nil {
		log.Fatal("Error fetching tables: ", err)
	}
}
