package db

import (
	"context"
	"database/sql"
	"log"
)

func CloseRows(rows *sql.Rows) {
	if err := rows.Close(); err != nil {
		log.Fatal(err)
	}
}

func CloseStatement(stmt *sql.Stmt) {
	if err := stmt.Close(); err != nil {
		log.Fatal(err)
	}
}

func ErrInRows(rows *sql.Rows) error {
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}

func MakeStatement(mssql *SQLServer, ctx context.Context, query string) *sql.Stmt {
	stmt, err := mssql.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	return stmt
}
