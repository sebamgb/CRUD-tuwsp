package db

import (
	"context"
	"fmt"
	"log"

	"tuwsp/models"
)

// UpdateProtocols update the protocol given from protocols by its id
func (mssql *SQLServer) UpdateProtocols(ctx context.Context, protocol *models.Protocol) error {
	// preparing statement
	query := `UPDATE url.protocols
	SET protocol = @p1
	WHERE id = @p2`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// update
	result, err := stmt.
		Exec(protocol.Protocol,
			protocol.Id)
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

// UpdateURLs update the domain and protocol_id given from urls by its id
func (mssql *SQLServer) UpdateURLs(ctx context.Context, url *models.Url) error {
	// preparing statement
	query := `UPDATE url.urls
	SET domain = @p1, protocol_id = @p2
	WHERE id = @p3`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// update
	result, err := stmt.
		Exec(url.Domain,
			url.ProtocolId,
			url.Id)
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

// UpdateEndpoints update the endpoint and url_id from endpoints by its id
func (mssql *SQLServer) UpdateEndpoints(ctx context.Context, endpoint *models.Endpoint) error {
	// preparing statement
	query := `UPDATE url.endpoints
	SET endpoint = @p1, url_id = @p2
	WHERE id = @p3`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// update
	result, err := stmt.
		Exec(endpoint.Endpoint,
			endpoint.UrlId,
			endpoint.Id)
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

// UpdateQueryKeys update the keyParam, urlId from query_keys by its id
func (mssql *SQLServer) UpdateQueryKeys(ctx context.Context, querykey *models.QueryKey) error {
	// preparing statement
	query := `UPDATE url.query_keys
	SET key_param = @p1, url_id = @p2
	WHERE id = @p3`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// update
	result, err := stmt.
		Exec(
			querykey.KeyParam,
			querykey.UrlId,
			querykey.Id)
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

// UpdateQueryValues update the valueParam, urlId from query_values by its id
func (mssql *SQLServer) UpdateQueryValues(ctx context.Context, queryvalue *models.QueryValue) error {
	// preparing statement
	query := `UPDATE url.query_values
	SET value_param = @p1, user_id = @p2
	WHERE id = @p3`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// update
	result, err := stmt.
		Exec(
			queryvalue.ValueParam,
			queryvalue.UserId,
			queryvalue.Id)
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

// UpdateUsers update the name, urlId given from users by its id
func (mssql *SQLServer) UpdateUsers(ctx context.Context, user *models.User) error {
	// preparing statement
	query := `UPDATE url.users
	SET name = @p1, url_id = @p2
	WHERE id = @p3`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// update
	result, err := stmt.
		Exec(
			user.Name,
			user.UrlId,
			user.Id)
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

// UpdateInfoUsers update phone, birthday given from info_users by its id
func (mssql *SQLServer) UpdateInfoUsers(ctx context.Context, infouser *models.InfoUser) error {
	// preparing statement
	query := `UPDATE url.info_users
	SET phone = @p1, birthday = @p2
	WHERE id = @p3`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// update
	result, err := stmt.
		Exec(infouser.Phone,
			infouser.Birthday,
			infouser.Id)
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
