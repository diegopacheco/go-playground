package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/machinebox/graphql"
)

func main() {

	client := graphql.NewClient("http://localhost:8080/graphql")

	// make a request
	req := graphql.NewRequest(`
	query ($key:ID!){
		bookById (id:$key) {
			name,
			pageCount,
			author {
			  id,
			  firstName,
			  lastName
			}
		}	
    }
	`)

	// set any variables
	req.Var("key", "book-3")

	// set header fields
	req.Header.Set("Cache-Control", "no-cache")

	// define a Context for the request
	ctx := context.Background()

	// run it and capture the response
	var respData map[string]interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}

	// Pretty print the json result
	b, err := json.MarshalIndent(respData, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))
}
