package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type MySqlClient struct {
	*sql.DB
}

func NewSqlClient(source string) *MySqlClient {
	db, err := sql.Open("mysql", source)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
        panic(err)
	}

	return &MySqlClient{db}
}

func (c *MySqlClient) ViewStats() sql.DBStats{
	return c.Stats()
}
