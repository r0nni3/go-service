package models

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type ConnectionObj struct {
	Host     string
	User     string
	Password string
	Database string
}

func GetConString() ConnectionObj {
	user, pass, database, host, port := os.Getenv("APP_USER"),
		os.Getenv("APP_PASS"), os.Getenv("APP_DB"),
		os.Getenv("APP_HOST"), os.Getenv("APP_PORT")

	if user == "" {
		user = "<DB-USER>"
	}
	if pass == "" {
		pass = "<DB-PASS>"
	}
	if database == "" {
		database = "<DB-NAME>"
	}
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "3306"
	}

	return ConnectionObj{
		User:     user,
		Password: pass,
		Database: database,
		Host:     host + ":" + port}
}

func (conObj ConnectionObj) OpenConnection() (*sql.DB, error) {
	conString := conObj.User + ":" + conObj.Password +
		"@tcp(" + conObj.Host + ")/" + conObj.Database
	con, err := sql.Open("mysql", conString)
	if err != nil {
		return nil, err
	}
	return con, nil
}
