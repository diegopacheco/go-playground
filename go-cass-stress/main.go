package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gocql/gocql"
)

// Keyspace Is the Cassandra Keyspace
const Keyspace = "CREATE KEYSPACE IF NOT EXISTS stress WITH REPLICATION = " +
	" { 'class' : 'SimpleStrategy', 'replication_factor' : 3 }; "

// Table is the Table schema for cass
const Table = "CREATE TABLE IF NOT EXISTS stress.TEST ( key text PRIMARY KEY, value text);"

// Insert is the Insert command for cass
const Insert = "INSERT INTO stress.TEST (key,value) VALUES (?, ?);"

func connect(host string) *gocql.Session {
	cluster := gocql.NewCluster(host)
	cluster.Consistency = gocql.Any
	session, err := cluster.CreateSession()
	if err != nil {
		fmt.Println("Could not create session in CASS. Err: ", err)
	}
	return session
}

func createKeyspace(session *gocql.Session) {
	if err := session.Query(Keyspace).Exec(); err != nil {
		fmt.Println("Could not create KEYSPACE in CASS. Err: ", err)
	}
	if err := session.Query(Table).Exec(); err != nil {
		fmt.Println("Could not create TABLE in CASS. Err: ", err)
	}
}

func insertRows(numRows int, session *gocql.Session) {
	errCount := 0
	okCount := 0
	for i := 0; i < numRows; i++ {
		if err := session.Query(Insert, "K"+strconv.Itoa(i), "V"+strconv.Itoa(i)).Exec(); err != nil {
			fmt.Println("Could not create INSERT in CASS. Err: ", err)
			errCount++
		} else {
			okCount++
		}
	}
	fmt.Printf(" DONE. %d rows inserted, %d errors. \n", okCount, errCount)
}

func parseArgs() (string, int) {
	rows := flag.Int("rows", 0, "Number of Records that you want to create in CASS. ")
	host := flag.String("host", "127.0.0.1", "CASS host(ip) you want connect. ")
	flag.Parse()

	if *rows <= 0 {
		fmt.Println("You need pass a possitive number of ROWS. i.e: -rows=100 ")
		os.Exit(1)
	}
	fmt.Printf(" %d recods will be created in CASS. \n", *rows)
	fmt.Printf(" Connecting to %s CASS. \n", *host)

	return *host, *rows
}

func main() {
	start := time.Now()
	host, rows := parseArgs()
	session := connect(host)
	createKeyspace(session)
	insertRows(rows, session)
	defer session.Close()

	elapsed := time.Since(start)
	fmt.Printf(" go-cass-stress took %s to insert %d rows.\n", elapsed, rows)
	fmt.Printf("FIN.\n")
}
