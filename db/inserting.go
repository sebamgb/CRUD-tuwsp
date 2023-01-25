package db

import (
	"context"
	"fmt"

	"tuwsp/models"
)

// Inserting userAuth
func (mssql *SQLServer) InsertIntoAuths(ctx context.Context, auth *models.Auth) error {
	// preparing statement
	query := `INSERT INTO info.auths (id, email, created_at, password)
	SELECT @p1, @p2, @p3, @p4
	WHERE NOT EXISTS(SELECT 1 FROM info.auths WHERE id = @p5)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	_, err := stmt.
		Exec(auth.Id,
			auth.Email,
			auth.CratedAt,
			auth.Password, auth.Id)
	if err != nil {
		return err
	}
	return nil
}

// Inserting protocol
func (mssql *SQLServer) InsertIntoProtocols(ctx context.Context, protocol *models.Protocol) error {
	// preparing statement
	query := `INSERT INTO url.protocols (id, protocol)
	SELECT @p1, @p2
	WHERE NOT EXISTS(SELECT 1 FROM url.protocols WHERE id = @p3)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	_, err := stmt.
		Exec(protocol.Id,
			protocol.Protocol, protocol.Id)
	if err != nil {
		return err
	}
	return nil
}

// inserting url
func (mssql *SQLServer) InsertIntoURLs(ctx context.Context, url *models.Url) error {
	// preparing statement
	query := `INSERT INTO url.urls (id, domain, protocol_id)
	SELECT @p1, @p2, @p3
	WHERE NOT EXISTS(SELECT 1 FROM url.urls WHERE id = @p4)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	_, err := stmt.
		Exec(url.Id,
			url.Domain,
			url.ProtocolId, url.Id)
	if err != nil {
		return err
	}
	return nil
}

// inserting endpoint
func (mssql *SQLServer) InsertIntoEndpoints(ctx context.Context, endpoint *models.Endpoint) error {
	// preparing statement
	query := `INSERT INTO url.endpoints (number, id, endpoint, url_id)
	SELECT @p1, @p2, @p3, @p4
	WHERE NOT EXISTS(SELECT 1 FROM url.endpoints WHERE number = @p5 AND url_id = @p6)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	_, err := stmt.
		Exec(endpoint.Number,
			endpoint.Id,
			endpoint.Endpoint,
			endpoint.UrlId,
			endpoint.Number, endpoint.UrlId)
	if err != nil {
		return err
	}
	return nil
}

// Inserting queryKey
func (mssql *SQLServer) InsertIntoQueryKeys(ctx context.Context, querykey *models.QueryKey) error {
	// preparing statement
	query := `INSERT INTO url.query_keys (number, id, key_param, url_id)
	SELECT @p1, @p2, @p3, @p4
	WHERE NOT EXISTS(SELECT 1 FROM url.query_keys WHERE number = @p5 AND url_id = @p6)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	_, err := stmt.
		Exec(querykey.Number,
			querykey.Id,
			querykey.KeyParam,
			querykey.UrlId,
			querykey.Number, querykey.UrlId)
	if err != nil {
		return err
	}
	return nil
}

// inserting queryValue
func (mssql *SQLServer) InsertIntoQueryValues(ctx context.Context, queryvalue *models.QueryValue) error {
	// preparing statement
	query := `INSERT INTO url.query_values (number, id, value_param, user_id)
	SELECT @p1, @p2, @p3, @p4
	WHERE NOT EXISTS(SELECT 1 FROM url.query_values WHERE number = @p5 AND user_id = @p6)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	_, err := stmt.
		Exec(queryvalue.Number,
			queryvalue.Id,
			queryvalue.ValueParam,
			queryvalue.UserId,
			queryvalue.Number, queryvalue.UserId)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	return nil
}

// Inserting data users
func (mssql *SQLServer) InsertIntoUsers(ctx context.Context, user *models.User) error {
	// preparing statement
	query := `INSERT INTO url.users (id, name, nick_name, url_id)
	SELECT @p1, @p2, @p3, @p4
	WHERE NOT EXISTS(SELECT 1 FROM url.users WHERE id = @p5)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	_, err := stmt.
		Exec(user.Id,
			user.Name,
			user.NickName,
			user.UrlId, user.Id)
	if err != nil {
		return err
	}
	return nil
}

// Inserting infoUser
func (mssql *SQLServer) InsertIntoInfoUsers(ctx context.Context, infouser *models.InfoUser) error {
	// preparing statement
	query := `INSERT INTO url.info_users (id, phone, country, cod_country, birthday, user_id)
	SELECT @p1, @p2, @p3, @p4, @p5, @p6
	WHERE NOT EXISTS(SELECT 1 FROM url.info_users WHERE id = @p7)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	_, err := stmt.
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
	return nil
}
