package repository

import (
	"context"

	"tuwsp/models"
)

type Repository interface {
	// Gets
	GetAuthById(ctx context.Context, id string) (*models.Auth, error)
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetInfoUserByIdAndUserId(ctx context.Context, id string, user_id string) (*models.InfoUser, error)
	GetProtocolById(ctx context.Context, id string) (*models.Protocol, error)
	GetUrlById(ctx context.Context, id string) (*models.Url, error)
	GetQueryKeyByUrlId(ctx context.Context, url_id string) ([]*models.QueryKey, error)
	GetQueryValueByUrlId(ctx context.Context, url_id string) ([]*models.QueryValue, error)
	GetEndPointByUrlId(ctx context.Context, url_id string) ([]*models.Endpoint, error)
	// Inserts
	InsertIntoAuths(ctx context.Context, auth *models.Auth) error
	InsertIntoProtocols(ctx context.Context, protocol *models.Protocol) error
	InsertIntoURLs(ctx context.Context, url *models.Url) error
	InsertIntoEndpoints(ctx context.Context, endpoint *models.Endpoint) error
	InsertIntoQueryKeys(ctx context.Context, querykey *models.QueryKey) error
	InsertIntoQueryValues(ctx context.Context, queryvalue *models.QueryValue) error
	InsertIntoUsers(ctx context.Context, user *models.User) error
	InsertIntoInfoUsers(ctx context.Context, infouser *models.InfoUser) error
	// Deletes
	DeleteProtocols(ctx context.Context, protocol *models.Protocol) error
	DeleteURLs(ctx context.Context, url *models.Url) error
	DeleteEndpoints(ctx context.Context, endpoint *models.Endpoint) error
	DeleteQueryKeys(ctx context.Context, querykey *models.QueryKey) error
	DeleteQueryValues(ctx context.Context, queryvalue *models.QueryValue) error
	DeleteUsers(ctx context.Context, user *models.User) error
	DeleteInfoUsers(ctx context.Context, infouser *models.InfoUser) error
	// Update
	UpdateProtocols(ctx context.Context, protocol *models.Protocol) error
	UpdateURLs(ctx context.Context, url *models.Url) error
	UpdateEndpoints(ctx context.Context, endpoint *models.Endpoint) error
	UpdateQueryKeys(ctx context.Context, querykey *models.QueryKey) error
	UpdateQueryValues(ctx context.Context, queryValue *models.QueryValue) error
	UpdateUsers(ctx context.Context, user *models.User) error
	UpdateInfoUsers(ctx context.Context, infouser *models.InfoUser) error
	// Cierre
	Close() error
}

var implementation Repository

// SetRepository assign a repository to implementation
func SetRepository(repository Repository) {
	implementation = repository
}

// GetAuthById do wich the implementation
func GetAuthById(ctx context.Context, id string) (*models.Auth, error) {
	return implementation.GetAuthById(ctx, id)
}

// GetUserById do wich the implementation
func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

// GetInfoUserByIdAndUserId do which the implementation
func GetInfoUserByIdAndUserId(ctx context.Context, id string, user_id string) (*models.InfoUser, error) {
	return implementation.GetInfoUserByIdAndUserId(ctx, id, user_id)
}

// GetProtocolById do which the implementation
func GetProtocolById(ctx context.Context, id string) (*models.Protocol, error) {
	return implementation.GetProtocolById(ctx, id)
}

// GetUrlById do which the implementation
func GetUrlById(ctx context.Context, id string) (*models.Url, error) {
	return implementation.GetUrlById(ctx, id)
}

// GetQueryKeyByUrlId do which the implementation
func GetQueryKeyByUrlId(ctx context.Context, url_id string) ([]*models.QueryKey, error) {
	return implementation.GetQueryKeyByUrlId(ctx, url_id)
}

// GetQueryValueByUrlId do which the implementation
func GetQueryValueByUrlId(ctx context.Context, url_id string) ([]*models.QueryValue, error) {
	return implementation.GetQueryValueByUrlId(ctx, url_id)
}

// GetEndPointByUrlId do which the implementation
func GetEndPointByUrlId(ctx context.Context, url_id string) ([]*models.Endpoint, error) {
	return implementation.GetEndPointByUrlId(ctx, url_id)
}

// InsertIntoAuths do which the implementation
func InsertIntoAuths(ctx context.Context, auth *models.Auth) error {
	return implementation.InsertIntoAuths(ctx, auth)
}

// InsertIntoProtocols do which the implementation
func InsertIntoProtocols(ctx context.Context, protocol *models.Protocol) error {
	return implementation.InsertIntoProtocols(ctx, protocol)
}

// InsertIntoURLs do which the implementation
func InsertIntoURLs(ctx context.Context, url *models.Url) error {
	return implementation.InsertIntoURLs(ctx, url)
}

// InsertIntoEndpoints do which the implementation
func InsertIntoEndpoints(ctx context.Context, endpoint *models.Endpoint) error {
	return implementation.InsertIntoEndpoints(ctx, endpoint)
}

// InsertIntoQueryKeys do which the implementation
func InsertIntoQueryKeys(ctx context.Context, querykey *models.QueryKey) error {
	return implementation.InsertIntoQueryKeys(ctx, querykey)
}

// InsertIntoQueryValues do which the implementation
func InsertIntoQueryValues(ctx context.Context, queryvalue *models.QueryValue) error {
	return implementation.InsertIntoQueryValues(ctx, queryvalue)
}

// InsertIntoUsers do which the implementation
func InsertIntoUsers(ctx context.Context, user *models.User) error {
	return implementation.InsertIntoUsers(ctx, user)
}

// InsertIntoInfoUsers do which the implementation
func InsertIntoInfoUsers(ctx context.Context, infouser *models.InfoUser) error {
	return implementation.InsertIntoInfoUsers(ctx, infouser)
}

// DeleteProtocols do which the implementation
func DeleteProtocols(ctx context.Context, protocol *models.Protocol) error {
	return implementation.DeleteProtocols(ctx, protocol)
}

// DeleteURLs do which the implementation
func DeleteURLs(ctx context.Context, url *models.Url) error {
	return implementation.DeleteURLs(ctx, url)
}

// DeleteEndpoints do which the implementation
func DeleteEndpoints(ctx context.Context, endpoint *models.Endpoint) error {
	return implementation.DeleteEndpoints(ctx, endpoint)
}

// DeleteQueryKeys do which the implementation
func DeleteQueryKeys(ctx context.Context, querykey *models.QueryKey) error {
	return implementation.DeleteQueryKeys(ctx, querykey)
}

// DeleteQueryValues do which the implementation
func DeleteQueryValues(ctx context.Context, queryvalue *models.QueryValue) error {
	return implementation.DeleteQueryValues(ctx, queryvalue)
}

// DeleteUsers do which the implementation
func DeleteUsers(ctx context.Context, user *models.User) error {
	return implementation.DeleteUsers(ctx, user)
}

// DeleteInfoUsers do which the implementation
func DeleteInfoUsers(ctx context.Context, infouser *models.InfoUser) error {
	return implementation.DeleteInfoUsers(ctx, infouser)
}

// UpdateProtocols do which the implementation
func UpdateProtocols(ctx context.Context, protocol *models.Protocol) error {
	return implementation.UpdateProtocols(ctx, protocol)
}

// UpdateURLs do which the implementation
func UpdateURLs(ctx context.Context, url *models.Url) error {
	return implementation.UpdateURLs(ctx, url)
}

// UpdateURLs do which the implementation
func UpdateEndpoints(ctx context.Context, endpoint *models.Endpoint) error {
	return implementation.UpdateEndpoints(ctx, endpoint)
}

// UpdateQueryKeys do which the implementation
func UpdateQueryKeys(ctx context.Context, querykey *models.QueryKey) error {
	return implementation.UpdateQueryKeys(ctx, querykey)
}

// UpdateQueryValues do which the implementation
func UpdateQueryValues(ctx context.Context, queryvalue *models.QueryValue) error {
	return implementation.UpdateQueryValues(ctx, queryvalue)
}

// UpdateUsers do which the implementation
func UpdateUsers(ctx context.Context, user *models.User) error {
	return implementation.UpdateUsers(ctx, user)
}

// UpdateProtocols do which the implementation
func UpdateInfoUsers(ctx context.Context, infouser *models.InfoUser) error {
	return implementation.UpdateInfoUsers(ctx, infouser)
}

// Close do which the implementation
func Close() error {
	return implementation.Close()
}
