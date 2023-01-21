package db

import (
	"context"
	"fmt"
	"log"

	"tuwsp/models"
)

// Inserting userAuth
func (mssql *SQLServer) InsertIntoAuths(ctx context.Context, auth *models.Auth) error {
	// preparing statement
	query := `INSERT INTO tuwsp_info.auths (id, email, created_at, password)
	SELECT @p1, @p2, @p3, @p4
	WHERE NOT EXISTS(SELECT 1 FROM auths WHERE id = @p5)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	result, err := stmt.
		Exec(auth.Id,
			auth.Email,
			auth.CratedAt,
			auth.Password, auth.Id)
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

// Inserting protocol
func (mssql *SQLServer) InsertIntoProtocols(ctx context.Context, protocol *models.Protocol) error {
	// preparing statement
	query := `INSERT INTO protocols (id, protocol)
	SELECT @p1, @p2
	WHERE NOT EXISTS(SELECT 1 FROM protocols WHERE id = @p3)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	result, err := stmt.
		Exec(protocol.Id,
			protocol.Protocol, protocol.Id)
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

// inserting url
func (mssql *SQLServer) InsertIntoURLs(ctx context.Context, url *models.Url) error {
	// preparing statement
	query := `INSERT INTO urls (id, domain, protocol_id)
	SELECT @p1, @p2, @p3
	WHERE NOT EXISTS(SELECT 1 FROM urls WHERE id = @p4)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	result, err := stmt.
		Exec(url.Id,
			url.Domain,
			url.ProtocolId, url.Id)
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

// inserting endpoint
func (mssql *SQLServer) InsertIntoEndpoints(ctx context.Context, endpoint *models.Endpoint) error {
	// preparing statement
	query := `INSERT INTO endpoints (id, endpoint, url_id)
	SELECT @p1, @p2, @p3
	WHERE NOT EXISTS(SELECT 1 FROM endpoints WHERE id = @p4)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	result, err := stmt.
		Exec(endpoint.Id,
			endpoint.Endpoint,
			endpoint.UrlId, endpoint.Id)
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

// Inserting queryKey
func (mssql *SQLServer) InsertIntoQueryKeys(ctx context.Context, querykey *models.QueryKey) error {
	// preparing statement
	query := `INSERT INTO query_keys (id, key_param, url_id)
	SELECT @p1, @p2, @p3
	WHERE NOT EXISTS(SELECT 1 FROM query_keys WHERE id = @p4)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	result, err := stmt.
		Exec(querykey.Id,
			querykey.KeyParam,
			querykey.UrlId, querykey.Id)
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

// inserting queryValue
func (mssql *SQLServer) InsertIntoQueryValues(ctx context.Context, queryvalue *models.QueryValue) error {
	// preparing statement
	query := `INSERT INTO query_values (id, value_param, url_id)
	SELECT @p1, @p2, @p3
	WHERE NOT EXISTS(SELECT 1 FROM query_values WHERE id = @p4)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	result, err := stmt.
		Exec(queryvalue.Id,
			queryvalue.ValueParam,
			queryvalue.UrlId, queryvalue.Id)
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

// Inserting data users
func (mssql *SQLServer) InsertIntoUsers(ctx context.Context, user *models.User) error {
	// preparing statement
	query := `INSERT INTO users (id, name, nick_name, url_id)
	SELECT @p1, @p2, @p3, @p4
	WHERE NOT EXISTS(SELECT 1 FROM users WHERE id = @p5)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	result, err := stmt.
		Exec(user.Id,
			user.Name,
			user.NickName,
			user.UrlId, user.Id)
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

// Inserting infoUser
func (mssql *SQLServer) InsertIntoInfoUsers(ctx context.Context, infouser *models.InfoUser) error {
	// preparing statement
	query := `INSERT INTO info_users (id, phone, country, cod_country, birthday, user_id)
	SELECT @p1, @p2, @p3, @p4, @p5, @p6
	WHERE NOT EXISTS(SELECT 1 FROM info_users WHERE id = @p7)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	result, err := stmt.
		Exec(
			infouser.Id,
			infouser.Phone,
			infouser.Country,
			infouser.CodCountry,
			infouser.Birthday,
			infouser.UserId, infouser.Id)
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
