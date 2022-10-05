/**
 *
 * @author nghiatc
 * @since Oct 05, 2022
 */

package main

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// Run code: go run main.go
// https://github.com/neo4j/neo4j-go-driver
func main() {
	uri := "neo4j://localhost:7687"
	username := "neo4j"
	password := "localtest"

	greeting, err := helloWorld(uri, username, password)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("greeting:", greeting)
}

func helloWorld(uri, username, password string) (string, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return "", err
	}
	defer driver.Close()

	// Sessions are short-lived, cheap to create and NOT thread safe.
	// Typically create one or more sessions per request in your web application.
	// Make sure to call Close on the session when done.
	// For multi-database support, set sessionConfig.DatabaseName to requested database
	// Session config will default to write mode,
	// if only reads are to be used configure session for read mode.
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	greeting, err := session.WriteTransaction(func(transaction neo4j.Transaction) (any, error) {
		result, err := transaction.Run(
			"CREATE (a:Greeting) SET a.message = $message RETURN a.message + ', from node ' + id(a)",
			map[string]any{"message": "hello, world"})
		if err != nil {
			return nil, err
		}
		if result.Next() {
			return result.Record().Values[0], nil
		}
		return nil, result.Err()
	})
	if err != nil {
		return "", err
	}
	return greeting.(string), nil
}
