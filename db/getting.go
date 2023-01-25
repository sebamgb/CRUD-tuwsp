package db

import (
	"context"
	"database/sql"
	"fmt"

	"tuwsp/models"
)

// GetAuthById get an auth from db by its id
func (mssql *SQLServer) GetAuthById(ctx context.Context, id string) (auth *models.Auth, err error) {
	// preparing statement
	query := `SELECT id, email, created_at, password
	FROM tuwsp_info.auths
	WHERE id = @p1`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// Query for a value based on a single row.
	if err = stmt.QueryRow(id).
		Scan(&auth.Id,
			&auth.Email,
			&auth.CratedAt,
			&auth.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("canPurchase %s: unknown user", id)
		}
		return nil, fmt.Errorf("canPurchase %s: %v", id, err)
	}
	return
}

// GetUserByNickName get an user from db by its id
func (mssql *SQLServer) GetUserByNickName(ctx context.Context, nick_name string) (user *models.User, err error) {
	// preparing statement
	query := `SELECT id, name, nick_name, url_id
	FROM users
	WHERE nick_name = @p1`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// Query for a value based on a single row.
	if err = stmt.QueryRow(nick_name).
		Scan(&user.Id,
			&user.Name,
			&user.NickName,
			&user.UrlId); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("canPurchase %s: unknown user", nick_name)
		}
		return nil, fmt.Errorf("canPurchase %s: %v", nick_name, err)
	}
	return
}

// GetInfoUserByUserId get multiple rows from db in info_Users by its id, user_id
func (mssql *SQLServer) GetInfoUserByUserId(ctx context.Context, user_id string) (infoUser *models.InfoUser, err error) {
	// preparing statement
	query := `SELECT id, phone, country, cod_country, birthday, user_id
	FROM info_users
	WHERE user_id = @p1`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// Query for a value based on a single row.
	if err := stmt.QueryRow(user_id).
		Scan(&infoUser.Id,
			&infoUser.Phone,
			&infoUser.Country,
			&infoUser.CodCountry,
			&infoUser.Birthday, &infoUser.UserId); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("canPurchase %s: unknown user", infoUser.Id)
		}
		return nil, fmt.Errorf("canPurchase %s: %v", infoUser.Id, err)
	}
	return
}

// GetProtocolById get a protocol from db by its id
func (mssql *SQLServer) GetProtocolById(ctx context.Context, id string) (protocol *models.Protocol, err error) {
	// preparing statement
	query := `SELECT id, protocol
	FROM protocols
	WHERE id = @p1`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// Query for a value based on a single row.
	//protocol = &models.Protocol{}
	if err := stmt.QueryRow(id).
		Scan(&protocol.Id,
			&protocol.Protocol); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("canPurchase %s: unknown protocol", id)
		}
		return nil, fmt.Errorf("canPurchase %s: %v", id, err)
	}
	return
}

// GetUrlById get an url from db by id
func (mssql *SQLServer) GetUrlById(ctx context.Context, id string) (_url *models.Url, err error) {
	// preparing statement
	query := `SELECT id, domain, protocol_id
	FROM urls
	WHERE id = @p1`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// Query for a value based on a single row.
	if err := stmt.QueryRow(id).
		Scan(
			&_url.Id,
			&_url.Domain,
			&_url.ProtocolId); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("canPurchase %s: unknown url", id)
		}
		return nil, fmt.Errorf("canPurchase %s: %v", id, err)
	}
	return
}

// GetQueryKeyByUrlId get keyParams by its urlId
func (mssql *SQLServer) GetQueryKeyByUrlId(ctx context.Context, url_id string) (queryKeys []*models.QueryKey, err error) {
	// preparing statement
	query := `SELECT id, key_param, url_id
	FROM query_keys
	WHERE url_id = @p1
	ORDER BY id asc`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	rows, err := stmt.Query(url_id)
	if err != nil {
		return nil, err
	}
	defer CloseRows(rows)
	for rows.Next() {
		var queryKey models.QueryKey
		if err = rows.
			Scan(&queryKey.Id,
				&queryKey.KeyParam,
				&queryKey.UrlId); err != nil {
			return
		}
		queryKeys = append(queryKeys, &queryKey)
	}
	err = ErrInRows(rows)
	return
}

// GetQueryValueByUserId get valueParams by its urlId
func (mssql *SQLServer) GetQueryValueByUserId(ctx context.Context, user_id string) (queryValues []*models.QueryValue, err error) {
	// preparing statement
	query := `SELECT id, key_param, user_id
	FROM query_values
	WHERE user_id = @p1
	ORDER BY id asc`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	rows, err := stmt.Query(user_id)
	if err != nil {
		return nil, err
	}
	defer CloseRows(rows)
	for rows.Next() {
		var queryValue models.QueryValue
		if err = rows.
			Scan(&queryValue.Id,
				&queryValue.ValueParam,
				&queryValue.UserId); err != nil {
			return
		}
		queryValues = append(queryValues, &queryValue)
	}
	err = ErrInRows(rows)
	return
}

// GetEndPointByUrlId get endpoints by its urlId
func (mssql *SQLServer) GetEndPointByUrlId(ctx context.Context, url_id string) (endpoints []*models.Endpoint, err error) {
	// preparing statement
	query := `SELECT id, endpoint, url_id
	FROM endpoints
	WHERE url_id = @p1
	ORDER BY id asc`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	rows, err := stmt.Query(url_id)
	if err != nil {
		return nil, err
	}
	defer CloseRows(rows)
	for rows.Next() {
		var endpoint models.Endpoint
		if err = rows.
			Scan(&endpoint.Id,
				&endpoint.Endpoint,
				&endpoint.UrlId); err != nil {
			return
		}
		endpoints = append(endpoints, &endpoint)
	}
	err = ErrInRows(rows)
	return
}

// Close cierra la conexión a la db
func (mssql *SQLServer) Close() error {
	return mssql.db.Close()
}
