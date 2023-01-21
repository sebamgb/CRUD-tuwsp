package db

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

type SQLServer struct {
	db *sql.DB
}

func (SQLServer) singleConnection(driver string, url string) (*sql.DB, error) {
	return sql.Open(driver, url)
}

// GetConnection get a connection pool
func GetSQLServer(driver string, url string) *SQLServer {
	var err error
	mssql := &SQLServer{}
	log.Println("Creating db")
	if mssql.db, err = mssql.singleConnection(driver, url); err != nil {
		log.Fatalf("imposible to connect: %v", err)
	}
	log.Println("Done")
	return mssql
}
