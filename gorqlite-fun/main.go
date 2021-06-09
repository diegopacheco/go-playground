package main

import (
	"fmt"

	"github.com/rqlite/gorqlite"
)

func main() {
	conn, _ := gorqlite.Open("http://localhost:4001/")
	conn.SetConsistencyLevel("strong")

	statementsCreate := make([]string, 0)
	sqlStmt := "create table secret_agents(id INTEGER NOT NULL PRIMARY KEY, hero_name TEXT, abbrev TEXT)"
	statementsCreate = append(statementsCreate, sqlStmt)
	_, err := conn.Write(statementsCreate)
	if err != nil {
		fmt.Printf("we have this error on creating agents table: %s\n", err.Error())
	}

	statements := make([]string, 0)
	pattern := "INSERT INTO secret_agents(id, hero_name, abbrev) VALUES (%d, '%s', '%3s')"
	statements = append(statements, fmt.Sprintf(pattern, 125718, "Speed Gibson", "Speed"))
	statements = append(statements, fmt.Sprintf(pattern, 209166, "Clint Barlow", "Clint"))
	statements = append(statements, fmt.Sprintf(pattern, 44107, "Barney Dunlap", "Barney"))
	results, _ := conn.Write(statements)
	for n, v := range results {
		fmt.Printf("for result %d, %d rows were affected\n", n, v.RowsAffected)
		if v.Err != nil {
			fmt.Printf("   we have this error: %s\n", v.Err.Error())
		}
	}

	var id int64
	var name string
	rows, err := conn.QueryOne("select id, hero_name from secret_agents where id > 500")
	if err != nil {
		fmt.Printf("we have this error on QUERY: %s\n", err.Error())
	}

	total := rows.NumRows()
	fmt.Printf("query returned %d rows\n", total)
	for rows.Next() {
		err = rows.Scan(&id, &name)
		fmt.Printf("*** %d %s \n", id, name)
		fmt.Printf("this is row number %d\n", rows.RowNumber()+1)
		fmt.Printf("there are %d rows overall %d \n", rows.RowNumber()+1, total)
	}
}
