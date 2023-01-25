package db

import (
	"context"
	"database/sql"
	"fmt"

	"tuwsp/models"
)

// GetDashboardByAuthId get a dashboard from db by its id
func (mssql *SQLServer) GetDashboardByAuthId(ctx context.Context, auth_id string) (*models.Dashboard, error) {
	// preparing statement
	query := `SELECT id, title, menu, app, owner
	FROM info.dashboards
	WHERE owner = @p1`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// Query for a value based on a single row.
	var dashboard = models.Dashboard{}
	if err := stmt.QueryRow(auth_id).
		Scan(&dashboard.Id,
			&dashboard.Title,
			&dashboard.Menu,
			&dashboard.App,
			&dashboard.Owner); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("errbyid:canPurchase %s: unknown user", auth_id)
		}
		return nil, fmt.Errorf("canPurchase %s: %v", auth_id, err)
	}
	return &dashboard, nil
}

// GetFormByTitle get a form from db by its title
func (mssql *SQLServer) GetFormByTitle(ctx context.Context, title string) (*models.Form, error) {
	// preparing statement
	query := `SELECT id, title, app, key_value
	FROM info.forms
	WHERE title = @p1`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// Query for a value based on a single row.
	var form = models.Form{}
	if err := stmt.QueryRow(title).
		Scan(&form.Id,
			&form.Title,
			&form.App,
			&form.Key_Value); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("errbyid:canPurchase %s: unknown user", title)
		}
		return nil, fmt.Errorf("canPurchase %s: %v", title, err)
	}
	return &form, nil
}

// GetLoginByAuthId get a login from db by its auth_id
func (mssql *SQLServer) GetLoginByAuthId(ctx context.Context, auth_id string) (*models.Login, error) {
	// preparing statement
	query := `SELECT id, email, password, created_at, log_out, auth_id, form_id
	FROM info.logins
	WHERE auth_id = @p1`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// Query for a value based on a single row.
	var login = models.Login{}
	if err := stmt.QueryRow(auth_id).
		Scan(&login.Id,
			&login.Email,
			&login.Password,
			&login.CreatedAt,
			&login.LogOut,
			&login.AuthId,
			&login.FormId); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("errbyid:canPurchase %s: unknown user", auth_id)
		}
		return nil, fmt.Errorf("canPurchase %s: %v", auth_id, err)
	}
	return &login, nil
}

// GetFormById get a form from db by its id
func (mssql *SQLServer) GetFormById(ctx context.Context, id string) (*models.Form, error) {
	// preparing statement
	query := `SELECT id, title, app, key_value
	FROM info.forms
	WHERE id = @p1`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// Query for a value based on a single row.
	var form = models.Form{}
	if err := stmt.QueryRow(id).
		Scan(&form.Id,
			&form.Title,
			&form.App,
			&form.Key_Value); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("errbyid:canPurchase %s: unknown user", id)
		}
		return nil, fmt.Errorf("canPurchase %s: %v", id, err)
	}
	return &form, nil
}

// GetLoginById get a login from db by its id
func (mssql *SQLServer) GetLoginById(ctx context.Context, id string) (*models.Login, error) {
	// preparing statement
	query := `SELECT id, email, password, created_at, log_out, auth_id, form_id
	FROM info.logins
	WHERE id = @p1`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// Query for a value based on a single row.
	var login = models.Login{}
	if err := stmt.QueryRow(id).
		Scan(&login.Id,
			&login.Email,
			&login.Password,
			&login.CreatedAt,
			&login.LogOut,
			&login.AuthId,
			&login.FormId); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("errbyid:canPurchase %s: unknown user", id)
		}
		return nil, fmt.Errorf("canPurchase %s: %v", id, err)
	}
	return &login, nil
}

// GetSignupById get a signup from db by its id
func (mssql *SQLServer) GetSignupById(ctx context.Context, id string) (*models.Signup, error) {
	// preparing statement
	query := `SELECT id, name, nick_name, email, password, confirm_password, form_id
	FROM info.signups
	WHERE id = @p1`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// Query for a value based on a single row.
	var signup = models.Signup{}
	if err := stmt.QueryRow(id).
		Scan(&signup.Id,
			&signup.Name,
			&signup.NickName,
			&signup.Email,
			&signup.Password,
			&signup.ConfirmPassword,
			&signup.FormId); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("errbyid:canPurchase %s: unknown user", id)
		}
		return nil, fmt.Errorf("canPurchase %s: %v", id, err)
	}
	return &signup, nil
}

// GetSignupByFormId get a signup from db by its form_id
func (mssql *SQLServer) GetSignupByFormId(ctx context.Context, form_id string) (*models.Signup, error) {
	// preparing statement
	query := `SELECT id, name, nick_name, email, password, confirm_password, form_id
	FROM info.signups
	WHERE form_id = @p1`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// Query for a value based on a single row.
	var signup = models.Signup{}
	if err := stmt.QueryRow(form_id).
		Scan(&signup.Id,
			&signup.Name,
			&signup.NickName,
			&signup.Email,
			&signup.Password,
			&signup.ConfirmPassword,
			&signup.FormId); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("errbyid:canPurchase %s: unknown user", form_id)
		}
		return nil, fmt.Errorf("canPurchase %s: %v", form_id, err)
	}
	return &signup, nil
}

// GetAuthById get an auth from db by its id
func (mssql *SQLServer) GetAuthById(ctx context.Context, id string) (*models.Auth, error) {
	// preparing statement
	query := `SELECT id, email, created_at, password
	FROM info.auths
	WHERE id = @p1`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// Query for a value based on a single row.
	var auth = models.Auth{}
	if err := stmt.QueryRow(id).
		Scan(&auth.Id,
			&auth.Email,
			&auth.CreatedAt,
			&auth.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("errbyid:canPurchase %s: unknown user", id)
		}
		return nil, fmt.Errorf("canPurchase %s: %v", id, err)
	}
	return &auth, nil
}

// GetAuthByEmail get an auth from db by its email
func (mssql *SQLServer) GetAuthByEmail(ctx context.Context, email string) (*models.Auth, error) {
	// preparing statement
	query := `SELECT id, email, created_at, password
	FROM info.auths
	WHERE email = @p1`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// Query for a value based on a single row.
	var auth = models.Auth{}
	if err := stmt.QueryRow(email).
		Scan(&auth.Id,
			&auth.Email,
			&auth.CreatedAt,
			&auth.Password); err != sql.ErrNoRows {
		return nil, err
	}
	return &auth, nil
}

// GetUserByNickName get an user from db by its nick_name
func (mssql *SQLServer) GetUserByNickName(ctx context.Context, nick_name string) (*models.User, error) {
	// preparing statement
	query := `SELECT id, name, nick_name, url_id
	FROM url.users
	WHERE nick_name = @p1`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// Query for a value based on a single row.
	var user = models.User{}
	if err := stmt.QueryRow(nick_name).
		Scan(&user.Id,
			&user.Name,
			&user.NickName,
			&user.UrlId); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("canPurchase %s: unknown user", nick_name)
		}
		return nil, fmt.Errorf("canPurchase %s: %v", nick_name, err)
	}
	return &user, nil
}

// GetUserById get an user from db by its id
func (mssql *SQLServer) GetUserById(ctx context.Context, id string) (*models.User, error) {
	// preparing statement
	query := `SELECT id, name, nick_name, url_id
	FROM url.users
	WHERE id = @p1`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// Query for a value based on a single row.
	var user = models.User{}
	if err := stmt.QueryRow(id).
		Scan(&user.Id,
			&user.Name,
			&user.NickName,
			&user.UrlId); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("canPurchase %s: unknown user id", id)
		}
		return nil, fmt.Errorf("canPurchase %s: %v", id, err)
	}
	return &user, nil
}

// GetInfoUserByUserId get multiple rows from db in info_Users by its user_id
func (mssql *SQLServer) GetInfoUserByUserId(ctx context.Context, user_id string) (*models.InfoUser, error) {
	// preparing statement
	query := `SELECT id, phone, country, cod_country, birthday, user_id
	FROM url.info_users
	WHERE user_id = @p1`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// Query for a value based on a single row.
	var infoUser = models.InfoUser{}
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
	return &infoUser, nil
}

// GetInfoUserByPhone get multiple rows from db in info_Users by its phone
func (mssql *SQLServer) GetInfoUserByPhone(ctx context.Context, phone int) (*models.InfoUser, error) {
	// preparing statement
	query := `SELECT id, phone, country, cod_country, birthday, user_id
	FROM url.info_users
	WHERE phone = @p1`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// Query for a value based on a single row.
	var infoUser = models.InfoUser{}
	if err := stmt.QueryRow(phone).
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
	return &infoUser, nil
}

// GetProtocolById get a protocol from db by its id
func (mssql *SQLServer) GetProtocolById(ctx context.Context, id string) (*models.Protocol, error) {
	// preparing statement
	query := `SELECT id, protocol
	FROM url.protocols
	WHERE id = @p1`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// Query for a value based on a single row.
	var protocol = models.Protocol{}
	if err := stmt.QueryRow(id).
		Scan(&protocol.Id,
			&protocol.Protocol); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("canPurchase %s: unknown protocol", id)
		}
		return nil, fmt.Errorf("canPurchase %s: %v", id, err)
	}
	return &protocol, nil
}

// GetUrlById get an url from db by id
func (mssql *SQLServer) GetUrlById(ctx context.Context, id string) (*models.Url, error) {
	// preparing statement
	query := `SELECT id, domain, protocol_id
	FROM url.urls
	WHERE id = @p1`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// Query for a value based on a single row.
	var _url = models.Url{}
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
	return &_url, nil
}

// GetQueryKeyByUrlId get keyParams by its urlId
func (mssql *SQLServer) GetQueryKeyByUrlId(ctx context.Context, url_id string) (queryKeys []*models.QueryKey, err error) {
	// preparing statement
	query := `SELECT number, id, key_param, url_id
	FROM url.query_keys
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
			Scan(&queryKey.Number,
				&queryKey.Id,
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
	query := `SELECT number, id, value_param, user_id
	FROM url.query_values
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
			Scan(&queryValue.Number,
				&queryValue.Id,
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
	FROM url.endpoints
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
