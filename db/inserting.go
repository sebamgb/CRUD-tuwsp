package db

import (
	"context"
	"fmt"

	"tuwsp/models"
)

// Inserting key value labels
func (mssql *SQLServer) InsertIntoKeyValueLabels(ctx context.Context, key_value *models.KeyValue) error {
	// preparing statement
	query := `INSERT INTO info.key_value_labels (id, title, label_name, label_nickname, label_email, label_phone, label_birthday, label_country, label_password, label_confirm_password, input_submit)
	SELECT @p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10, @p11
	WHERE NOT EXISTS(SELECT 1 FROM info.key_value_labels WHERE id = @p12)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	_, err := stmt.
		Exec(key_value.Id, key_value.Title,
			key_value.LabelName, key_value.LabelNickname,
			key_value.LabelEmail, key_value.LabelPhone,
			key_value.LabelBirthday, key_value.LabelCountry,
			key_value.LabelPassword, key_value.LabelConfirmPassword,
			key_value.InputSubmit, key_value.Id)
	if err != nil {
		return err
	}
	return nil
}

// Inserting key value placeholders
func (mssql *SQLServer) InsertIntoKeyValuePlaceholders(ctx context.Context, key_value *models.KeyValue) error {
	// preparing statement
	query := `INSERT INTO info.key_value_placeholders (id, placeholder_name, placeholder_nickname, placeholder_email, placeholder_phone, placeholder_country, placeholder_password, placeholder_confirm_password, label_id)
	SELECT @p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9
	WHERE NOT EXISTS(SELECT 1 FROM info.key_value_labels WHERE id = @p10)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	_, err := stmt.
		Exec(key_value.Id, key_value.PlaceholderName,
			key_value.PlaceholderNickname, key_value.PlaceholderEmail,
			key_value.PlaceholderPhone, key_value.PlaceholderCountry,
			key_value.PlaceholderPassword, key_value.PlaceholderConfirmPassword,
			key_value.LabelId, key_value.Id)
	if err != nil {
		return err
	}
	return nil
}

// Inserting login
func (mssql *SQLServer) InsertIntoLogins(ctx context.Context, login *models.Login) error {
	// preparing statement
	query := `INSERT INTO info.logins (id, title, url, method, email, created_at, password, log_out, auth_id, form_id)
	SELECT @p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10
	WHERE NOT EXISTS(SELECT 1 FROM info.logins WHERE id = @p11)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	_, err := stmt.
		Exec(login.Id, login.Title,
			login.Url, login.Method,
			login.Email, login.CreatedAt,
			login.Password, login.LogOut,
			login.AuthId, login.FormId, login.Id)
	if err != nil {
		return err
	}
	return nil
}

// Inserting dashboard
func (mssql *SQLServer) InsertIntoDashboards(ctx context.Context, dashboard *models.Dashboard) error {
	// preparing statement
	query := `INSERT INTO info.dashboards (id, title, menu, app, owner)
	SELECT @p1, @p2, @p3, @p4, @p5
	WHERE NOT EXISTS(SELECT 1 FROM info.dashboards WHERE id = @p6)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	_, err := stmt.
		Exec(dashboard.Id,
			dashboard.Title,
			dashboard.Menu,
			dashboard.App,
			dashboard.Owner, dashboard.Id)
	if err != nil {
		return err
	}
	return nil
}

// Inserting form
func (mssql *SQLServer) InsertIntoForms(ctx context.Context, form *models.Form) error {
	// preparing statement
	query := `INSERT INTO info.forms (id, app, key_value)
	SELECT @p1, @p2, @p3
	WHERE NOT EXISTS(SELECT 1 FROM info.forms WHERE id = @p4)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	_, err := stmt.
		Exec(form.Id,
			form.App,
			form.Key_Value, form.Id)
	if err != nil {
		return err
	}
	return nil
}

// Inserting signup
func (mssql *SQLServer) InsertIntoSignups(ctx context.Context, signup *models.Signup) error {
	// preparing statement
	query := `INSERT INTO info.signups (id, title, url, method, name, nick_name, email, phone, password, confirm_password, country, form_id)
	SELECT @p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10, @p11, @p12
	WHERE NOT EXISTS(SELECT 1 FROM info.signups WHERE id = @p13)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	_, err := stmt.
		Exec(signup.Id, signup.Title,
			signup.Url, signup.Method,
			signup.Name, signup.NickName,
			signup.Email, signup.Phone,
			signup.Password, signup.ConfirmPassword,
			signup.Country, signup.FormId,
			signup.Id)
	if err != nil {
		return err
	}
	return nil
}

// Inserting userAuth
func (mssql *SQLServer) InsertIntoAuths(ctx context.Context, auth *models.Auth) error {
	// preparing statement
	query := `INSERT INTO info.auths (id, email, created_at, password)
	SELECT @p1, @p2, @p3, @p4
	WHERE NOT EXISTS(SELECT 1 FROM info.auths WHERE id = @p5 AND email <> "" AND password <> "")`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	_, err := stmt.
		Exec(auth.Id,
			auth.Email,
			auth.CreatedAt,
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
