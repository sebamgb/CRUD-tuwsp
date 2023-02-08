package db

import (
	"context"
	"fmt"
	"log"
)

// Deleting login
func (mssql *SQLServer) DeleteLogins(ctx context.Context, id string) error {
	// preparing statement
	query := `DELETE FROM info.logins
	WHERE EXISTS (SELECT 1 FROM logins
	WHERE id = @p1)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// delete
	result, err := stmt.
		Exec(id)
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

// Deleting dashboard
func (mssql *SQLServer) DeleteDashboards(ctx context.Context, id string) error {
	// preparing statement
	query := `DELETE FROM info.dashboards
	WHERE EXISTS (SELECT 1 FROM dashboards
	WHERE id = @p1)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// delete
	result, err := stmt.
		Exec(id)
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

// Delete Form
func (mssql *SQLServer) DeleteForms(ctx context.Context, id string) error {
	// preparing statement
	query := `DELETE FROM info.forms
	WHERE EXISTS (SELECT 1 FROM forms
	WHERE id = @p1)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// delete
	result, err := stmt.
		Exec(id)
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

// Deleting auth
func (mssql *SQLServer) DeleteAuths(ctx context.Context, id string) error {
	// preparing statement
	query := `DELETE FROM info.auths
	WHERE EXISTS (SELECT 1 FROM auths
	WHERE id = @p1)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// delete
	result, err := stmt.
		Exec(id)
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

// Deleting signup
func (mssql *SQLServer) DeleteSignups(ctx context.Context, id string) error {
	// preparing statement
	query := `DELETE FROM info.signups
	WHERE EXISTS (SELECT 1 FROM signups
	WHERE id = @p1)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// delete
	result, err := stmt.
		Exec(id)
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

// Deleting protocol
func (mssql *SQLServer) DeleteProtocols(ctx context.Context, id string) error {
	// preparing statement
	query := `DELETE FROM url.protocols
	WHERE EXISTS (SELECT 1 FROM protocols
	WHERE id = @p1)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	result, err := stmt.
		Exec(id)
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
func (mssql *SQLServer) DeleteURLs(ctx context.Context, id string) error {
	// preparing statement
	query := `DELETE FROM url.urls
	WHERE EXISTS (SELECT 1 FROM urls
	WHERE id = @p1)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	result, err := stmt.
		Exec(id)
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
func (mssql *SQLServer) DeleteEndpoints(ctx context.Context, id string) error {
	// preparing statement
	query := `DELETE FROM url.endpoints
	WHERE EXISTS (SELECT 1 FROM endpoints
	WHERE id = @p1)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	result, err := stmt.
		Exec(id)
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
func (mssql *SQLServer) DeleteQueryKeys(ctx context.Context, id string) error {
	// preparing statement
	query := `DELETE FROM url.query_keys
	WHERE EXISTS (SELECT 1 FROM query_keys
	WHERE id = @p1)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	result, err := stmt.
		Exec(id)
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
func (mssql *SQLServer) DeleteQueryValues(ctx context.Context, id string) error {
	// preparing statement
	query := `DELETE FROM url.query_values
	WHERE EXISTS (SELECT 1 FROM query_values
	WHERE id = @p1)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	result, err := stmt.
		Exec(id)
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
func (mssql *SQLServer) DeleteUsers(ctx context.Context, id string) error {
	// preparing statement
	query := `DELETE FROM url.users
	WHERE EXISTS (SELECT 1 FROM users
	WHERE id = @p1)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	result, err := stmt.
		Exec(id)
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
func (mssql *SQLServer) DeleteInfoUsers(ctx context.Context, id string) error {
	// preparing statement
	query := `DELETE FROM url.info_users
	WHERE EXISTS (SELECT 1 FROM users
	WHERE id = @p1 AND phone = @p2 AND user_id = @p3)`
	stmt := MakeStatement(mssql, ctx, query)
	defer CloseStatement(stmt)
	// insert
	result, err := stmt.
		Exec(id)
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
