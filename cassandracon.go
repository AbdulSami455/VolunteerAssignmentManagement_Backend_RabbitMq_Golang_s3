package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func cluster_connection() {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "system_schema" // Use system_schema keyspace to query system tables
	cluster.Consistency = gocql.Quorum

	// Create a new session
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal("Error creating session: ", err)
	}
	defer session.Close()

	// Query to fetch all tables in the keyspace
	query := "SELECT table_name FROM tables WHERE keyspace_name = 'your_keyspace_name'"

	iter := session.Query(query).Iter()
	var tableName string

	fmt.Println("Tables in your_keyspace_name:")
	for iter.Scan(&tableName) {
		fmt.Println(tableName)
	}
	if err := iter.Close(); err != nil {
		log.Fatal("Error fetching tables: ", err)
	}
}
