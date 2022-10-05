/**
 *
 * @author nghiatc
 * @since Oct 06, 2022
 */

package ndb

import (
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

// NDBName var
var (
	NDBName = "neo4j"
)

var NClient neo4j.Driver

// InitNeo4j Init Neo4j client
// https://github.com/neo4j/neo4j-go-driver
func InitNeo4j() {
	// c := nconf.GetConfig()
	// uri := c.GetString("neo4jdb.uri")           // "neo4j://localhost:7687"
	// username := c.GetString("neo4jdb.username") // "neo4j"
	// password := c.GetString("neo4jdb.password") // "localtest"
	// NDBName = c.GetString("neo4jdb.name")       // "neo4j"
	uri := "neo4j://localhost:7687"
	username := "neo4j"
	password := "localtest"

	// Neo4j 4.0, defaults to no TLS therefore use bolt:// or neo4j://
	NClient, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Neo4j!")
	// nlogger.NLog.Logger.Info("Connected to Neo4j!")
}

func NClose() {
	if NClient != nil {
		NClient.Close()
	}
	fmt.Println("Disconnect Neo4j!")
	// nlogger.NLog.Logger.Info("Disconnect Neo4j!")
}

// NewSessionDefault Sessions are short-lived, cheap to create and NOT thread safe.
// Typically create one or more sessions per request in your web application.
// Make sure to call Close on the session when done.
// For multi-database support, set sessionConfig.DatabaseName to requested database
// Session config will default to write mode,
// if only reads are to be used configure session for read mode.
func NewSessionDefault(isWrite bool) neo4j.Session {
	if isWrite {
		session := NClient.NewSession(neo4j.SessionConfig{
			AccessMode:   neo4j.AccessModeWrite,
			DatabaseName: NDBName,
		})
		return session
	} else {
		session := NClient.NewSession(neo4j.SessionConfig{
			AccessMode:   neo4j.AccessModeRead,
			DatabaseName: NDBName,
		})
		return session
	}
}
