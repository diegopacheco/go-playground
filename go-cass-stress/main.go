package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/gocql/gocql"
)

const KEYSPACE_NAME = "stress"
const KEYSPACE = "CREATE KEYSPACE stress WITH REPLICATION = " +
	" { 'class' : 'SimpleStrategy', 'replication_factor' : 3 }; "
const TABLE = "CREATE TABLE stress.TEST ( key text PRIMARY KEY, value text);"
const INSERT = "INSERT INTO stress.TEST (key,value) VALUES ('?', '?');"

func connect(host string) *gocql.Session {
	cluster := gocql.NewCluster(host)
	session, err := cluster.CreateSession()
	if err != nil {
		fmt.Errorf("Could not create session in CASS. Err: ", err)
	}
	return session
}

func createKeyspace(session *gocql.Session) {
	if err := session.Query(KEYSPACE).Exec(); err != nil {
		fmt.Errorf("Could not create KEYSPACE in CASS. Err: ", err)
	}
	if err := session.Query(TABLE).Exec(); err != nil {
		fmt.Errorf("Could not create TABLE in CASS. Err: ", err)
	}
}

func insertRows(numRows int, session *gocql.Session) {
	for i := 0; i < numRows; i++ {
		if err := session.Query(INSERT, "K"+strconv.Itoa(i), "V"+strconv.Itoa(i)).Exec(); err != nil {
			fmt.Errorf("Could not create INSERT in CASS. Err: ", err)
		}
	}
}

func parseRowsArg() int {
	rows := flag.Int("rows", 0, "Number of Records that you want to create in CASS. ")
	flag.Parse()
	if *rows <= 0 {
		fmt.Println("You need pass a possitive number of ROWS. i.e: -rows=100 ")
		os.Exit(1)
	}
	fmt.Printf(" %d recods will be created in CASS. \n", *rows)
	return *rows
}

func parseHostArg() string {
	host := flag.String("host", "", "CASS host(ip) you want connect. ")
	flag.Parse()
	fmt.Printf(" Connecting to %s CASS. \n", *host)
	return *host
}

func main() {
	session := connect(parseHostArg())
	createKeyspace(session)
	insertRows(parseRowsArg(), session)
	defer session.Close()
}
