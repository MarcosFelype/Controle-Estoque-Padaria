package postgres

import (
	"errors"
	"fmt"
)

var (
	user     string
	password string
	dbName   string
	host     string
	port     int
)

func SetUpCredentials(newUser, newPwd, newDbName, newHost string, newPort int) error {
	user = newUser
	password = newPwd
	dbName = newDbName
	host = newHost
	port = newPort

	if !HasValidCredentials() {
		return errors.New("invalid credentials for the postgres database")
	}

	return nil
}

func HasValidCredentials() bool {
	hasTheRequiredFields := host != "" && dbName != "" && port > 0

	if password != "" {
		return hasTheRequiredFields && user != ""
	}

	return hasTheRequiredFields
}

func getPostgresConnectionURI() (string, error) {
	if !HasValidCredentials() {
		return "", errors.New("invalid credentials for the postgres database")
	}

	if dbName == "" {
		return fmt.Sprintf("postgres://%s:%d/%s?sslmode=disable", host, port, dbName), nil
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbName), nil
}
