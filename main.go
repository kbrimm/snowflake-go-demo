package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	sf "github.com/snowflakedb/gosnowflake"
)

// Parameters holds the contents of parameters.json
type Parameters struct {
	Account   string `json:"account"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Warehouse string `json:"warehouse"`
	Database  string `json:"database"`
	Schema    string `json:"schema"`
}

// getDSN constructs a DSN based on the test connection parameters
func getDSN() (string, *sf.Config, error) {
	paramConfig, err := os.Open("parameters.json")
	byteStream, _ := ioutil.ReadAll(paramConfig)
	var params Parameters
	json.Unmarshal(byteStream, &params)

	if err != nil {
		panic("unable to read parameters.json")
	}

	cfg := &sf.Config{
		Account:   params.Account,
		User:      params.User,
		Password:  params.Password,
		Warehouse: params.Warehouse,
		Database:  params.Database,
		Schema:    params.Schema,
	}

	dsn, err := sf.DSN(cfg)
	return dsn, cfg, err
}

func selectOne() {
	dsn, cfg, err := getDSN()
	if err != nil {
		log.Fatalf("failed to create DSN from Config: %v, err: %v", cfg, err)
	}

	db, err := sql.Open("snowflake", dsn)
	if err != nil {
		log.Fatalf("failed to connect. %v, err: %v", dsn, err)
	}
	defer db.Close()
	query := "SELECT 1 FROM TEST"
	rows, err := db.Query(query) // no cancel is allowed
	if err != nil {
		log.Fatalf("failed to run a query. %v, err: %v", query, err)
	}
	defer rows.Close()
	var v int
	for rows.Next() {
		err := rows.Scan(&v)
		if err != nil {
			log.Fatalf("failed to get result. err: %v", err)
		}
		if v != 1 {
			log.Fatalf("failed to get 1. got: %v", v)
		}
	}
	if rows.Err() != nil {
		fmt.Printf("ERROR: %v\n", rows.Err())
		return
	}
	fmt.Printf("Congrats! You have successfully run %v with Snowflake DB!\n", query)
}

func main() {
	fmt.Println("Running snowflake-go-demo")
	selectOne()
}
