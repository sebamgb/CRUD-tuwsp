package db

import (
	"context"
	"fmt"
	"log"

	"tuwsp/models"
)

// UpdateLoginsInLogIn update the email, password, auth_id given from logins by its id
func (mssql *SQLServer) UpdateLoginsInLogIn(ctx context.Context, login *models.Login) error {
	// preparing statement
	query := `UPDATE info.logins
	SET email = @p1, password = @p2, auth_id = @p3
	WHERE id = @p4`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// update
	result, err := stmt.
		Exec(login.Email,
			login.Password,
			login.AuthId,
			login.Id)
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

// UpdateLoginsInLogOut update the log_out given from logins by its id
func (mssql *SQLServer) UpdateLoginsInLogOut(ctx context.Context, login *models.Login) error {
	// preparing statement
	query := `UPDATE info.logins
	SET log_out = @p1
	WHERE id = @p2`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// update
	result, err := stmt.
		Exec(login.LogOut,
			login.Id)
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

// UpdateDashboards update the title, menu, app and owner given from dashboards by its id
func (mssql *SQLServer) UpdateDashboards(ctx context.Context, dashboard *models.Dashboard) error {
	// preparing statement
	query := `UPDATE info.dashboards
	SET title = @p1, menu = @p2, app = @p3, owner = @p4
	WHERE id = @p5`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// update
	result, err := stmt.
		Exec(dashboard.Title,
			dashboard.Menu,
			dashboard.App,
			dashboard.Owner,
			dashboard.Id)
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

// UpdateForms update app and key_value given from forms by its id
func (mssql *SQLServer) UpdateForms(ctx context.Context, form *models.Form) error {
	// preparing statement
	query := `UPDATE info.forms
	SET app = @p1, key_value = @p2
	WHERE id = @p3`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// update
	result, err := stmt.
		Exec(form.App,
			form.Key_Value,
			form.Id)
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

// UpdateAuths update the email, created_at, password and signup_id given from auths by its id
func (mssql *SQLServer) UpdateAuths(ctx context.Context, auth *models.Auth) error {
	// preparing statement
	query := `UPDATE info.auths
	SET email = @p1, created_at = @p2, password = @p3, signup_id = @p4
	WHERE id = @p5`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// update
	result, err := stmt.
		Exec(auth.Email,
			auth.CreatedAt,
			auth.Password,
			auth.SignupId,
			auth.Id)
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

// UpdateSignups update the name, nick_name, email, phone, password and form_id given from signups by its id
func (mssql *SQLServer) UpdateSignups(ctx context.Context, signup *models.Signup) error {
	// preparing statement
	query := `UPDATE info.signups
	SET name = @p1, nick_name = @p2, email = @p3, phone = @p4, password = @p5, form_id = @p6
	WHERE id = @p7`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// update
	result, err := stmt.
		Exec(signup.Name,
			signup.NickName,
			signup.Email,
			signup.Phone,
			signup.Password,
			signup.FormId,
			signup.Id)
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
