package db

import (
	"context"
	"fmt"
	"log"

	"tuwsp/models"
)

// Deleting protocol
func (mssql *SQLServer) DeleteProtocols(ctx context.Context, protocol *models.Protocol) error {
	// preparing statement
	query := `DELETE FROM protocols
	WHERE EXISTS (SELECT 1 FROM protocols
	WHERE id = @p1 AND protocol = @p2)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	result, err := stmt.
		Exec(
			protocol.Id,
			protocol.Protocol)
	if err != nil {
		return err
	}
	// rows affected?
	if rows, err := result.RowsAffected(); err != nil {
		return fmt.Errorf("error: %s when try to get rows affected",
			err.Error())
	} else {
		log.Printf("Rows Affected: %d", rows)
	}
	return nil
}

// Dleting url
func (mssql *SQLServer) DeleteURLs(ctx context.Context, url *models.Url) error {
	// preparing statement
	query := `DELETE FROM urls
	WHERE EXISTS (SELECT 1 FROM urls
	WHERE id = @p1 AND domain = @p2 AND protocol_id = @p3)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	result, err := stmt.
		Exec(
			url.Id,
			url.Domain,
			url.ProtocolId)
	if err != nil {
		return err
	}
	// rows affected?
	if rows, err := result.RowsAffected(); err != nil {
		return fmt.Errorf("error: %s when try to get rows affected",
			err.Error())
	} else {
		log.Printf("Rows Affected: %d", rows)
	}
	return nil
}

// Deleting endpoint
func (mssql *SQLServer) DeleteEndpoints(ctx context.Context, endpoint *models.Endpoint) error {
	// preparing statement
	query := `DELETE FROM endpoints
	WHERE EXISTS (SELECT 1 FROM endpoints
	WHERE id = @p1 AND endpoint = @p2 AND url_id = @p3)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	result, err := stmt.
		Exec(
			endpoint.Id,
			endpoint.Endpoint,
			endpoint.UrlId)
	if err != nil {
		return err
	}
	// rows affected?
	if rows, err := result.RowsAffected(); err != nil {
		return fmt.Errorf("error: %s when try to get rows affected",
			err.Error())
	} else {
		log.Printf("Rows Affected: %d", rows)
	}
	return nil
}

// Deleting queryKey
func (mssql *SQLServer) DeleteQueryKeys(ctx context.Context, querykey *models.QueryKey) error {
	// preparing statement
	query := `DELETE FROM query_keys
	WHERE EXISTS (SELECT 1 FROM query_keys
	WHERE id = @p1 AND key_param = @p2 AND url_id = @p3)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	result, err := stmt.
		Exec(
			querykey.Id,
			querykey.KeyParam,
			querykey.UrlId)
	if err != nil {
		return err
	}
	// rows affected?
	if rows, err := result.RowsAffected(); err != nil {
		return fmt.Errorf("error: %s when try to get rows affected",
			err.Error())
	} else {
		log.Printf("Rows Affected: %d", rows)
	}
	return nil
}

// Deleting queryValue
func (mssql *SQLServer) DeleteQueryValues(ctx context.Context, queryvalue *models.QueryValue) error {
	// preparing statement
	query := `DELETE FROM query_values
	WHERE EXISTS (SELECT 1 FROM query_values
	WHERE id = @p1 AND value_param = @p2 AND url_id = @p3)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	result, err := stmt.
		Exec(
			queryvalue.Id,
			queryvalue.ValueParam,
			queryvalue.UrlId)
	if err != nil {
		return err
	}
	// rows affected?
	if rows, err := result.RowsAffected(); err != nil {
		return fmt.Errorf("error: %s when try to get rows affected",
			err.Error())
	} else {
		log.Printf("Rows Affected: %d", rows)
	}
	return nil
}

// Deleting user
func (mssql *SQLServer) DeleteUsers(ctx context.Context, user *models.User) error {
	// preparing statement
	query := `DELETE FROM users
	WHERE EXISTS (SELECT 1 FROM users
	WHERE id = @p1 AND name = @p2 AND url_id = @p3)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	result, err := stmt.Exec(user.Id, user.Name, user.UrlId)
	if err != nil {
		return err
	}
	// rows affected?
	if rows, err := result.RowsAffected(); err != nil {
		return fmt.Errorf("error: %s when try to get rows affected",
			err.Error())
	} else {
		log.Printf("Rows Affected: %d", rows)
	}
	return nil
}

// Deleting infoUser
func (mssql *SQLServer) DeleteInfoUsers(ctx context.Context, infouser *models.InfoUser) error {
	// preparing statement
	query := `DELETE FROM info_users
	WHERE EXISTS (SELECT 1 FROM users
	WHERE id = @p1 AND phone = @p2 AND user_id = @p3)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	result, err := stmt.Exec(infouser.Id, infouser.Phone, infouser.UserId)
	if err != nil {
		return err
	}
	// rows affected?
	if rows, err := result.RowsAffected(); err != nil {
		return fmt.Errorf("error: %s when try to get rows affected",
			err.Error())
	} else {
		log.Printf("Rows Affected: %d", rows)
	}
	return nil
}
