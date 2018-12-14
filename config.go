package main

import (
	"fmt"
	"os"
)

func DbConnectionString() (string, error) {
	host := os.Getenv("POSTGRES_IP")
	dbName := "postgres"
	if host == "" {
		host = "localhost"
		dbName = "raceDb"
	}

	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		user = "foo"
	}

	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		password = "foo"
	}

	return fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", user, dbName, password, host), nil
}
