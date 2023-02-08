package repository

import (
	"context"

	"tuwsp/models"
)

type Repository interface {
	// Gets
	GetKeyValueById(ctx context.Context, id string) (*models.KeyValue, error)
	GetFormByTitle(ctx context.Context, title string) (*models.Form, error)
	GetFormById(ctx context.Context, id string) (*models.Form, error)
	GetLoginByAuthId(ctx context.Context, auth_id string) (*models.Login, error)
	GetSignupById(ctx context.Context, id string) (*models.Signup, error)
	GetSignupByFormId(ctx context.Context, form_id string) (*models.Signup, error)
	GetDashboardByAuthId(ctx context.Context, auth_id string) (*models.Dashboard, error)
	GetAuthById(ctx context.Context, id string) (*models.Auth, error)
	GetAuthByEmail(ctx context.Context, email string) (*models.Auth, error)
	GetUserByNickName(ctx context.Context, nick_name string) (*models.User, error)
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetInfoUserByUserId(ctx context.Context, user_id string) (*models.InfoUser, error)
	GetInfoUserByPhone(ctx context.Context, phone int) (*models.InfoUser, error)
	GetProtocolById(ctx context.Context, id string) (*models.Protocol, error)
	GetUrlById(ctx context.Context, id string) (*models.Url, error)
	GetQueryKeyByUrlId(ctx context.Context, url_id string) ([]*models.QueryKey, error)
	GetQueryValueByUserId(ctx context.Context, user_id string) ([]*models.QueryValue, error)
	GetEndPointByUrlId(ctx context.Context, url_id string) ([]*models.Endpoint, error)
	// lists
	ListForms(ctx context.Context) ([]*models.Form, error)
	ListLogins(ctx context.Context) ([]*models.Login, error)
	ListSignups(ctx context.Context) ([]*models.Signup, error)
	ListDashboards(ctx context.Context) ([]*models.Dashboard, error)
	ListAuths(ctx context.Context) ([]*models.Auth, error)
	ListProtocols(ctx context.Context) ([]*models.Protocol, error)
	ListURLs(ctx context.Context) ([]*models.Url, error)
	ListEndpoints(ctx context.Context) ([]*models.Endpoint, error)
	ListQueryKeys(ctx context.Context) ([]*models.QueryKey, error)
	ListQueryValues(ctx context.Context) ([]*models.QueryValue, error)
	ListUsers(ctx context.Context) ([]*models.User, error)
	ListInfoUsers(ctx context.Context) ([]*models.InfoUser, error)
	// Inserts
	InsertIntoSignups(ctx context.Context, signup *models.Signup) error
	InsertIntoLogins(ctx context.Context, login *models.Login) error
	InsertIntoForms(ctx context.Context, form *models.Form) error
	InsertIntoDashboards(ctx context.Context, dashboard *models.Dashboard) error
	InsertIntoAuths(ctx context.Context, auth *models.Auth) error
	InsertIntoProtocols(ctx context.Context, protocol *models.Protocol) error
	InsertIntoURLs(ctx context.Context, url *models.Url) error
	InsertIntoEndpoints(ctx context.Context, endpoint *models.Endpoint) error
	InsertIntoQueryKeys(ctx context.Context, querykey *models.QueryKey) error
	InsertIntoQueryValues(ctx context.Context, queryvalue *models.QueryValue) error
	InsertIntoUsers(ctx context.Context, user *models.User) error
	InsertIntoInfoUsers(ctx context.Context, infouser *models.InfoUser) error
	// Deletes
	DeleteLogins(ctx context.Context, id string) error
	DeleteSignups(ctx context.Context, id string) error
	DeleteForms(ctx context.Context, id string) error
	DeleteDashboards(ctx context.Context, id string) error
	DeleteAuths(ctx context.Context, id string) error
	DeleteProtocols(ctx context.Context, id string) error
	DeleteURLs(ctx context.Context, id string) error
	DeleteEndpoints(ctx context.Context, id string) error
	DeleteQueryKeys(ctx context.Context, id string) error
	DeleteQueryValues(ctx context.Context, id string) error
	DeleteUsers(ctx context.Context, id string) error
	DeleteInfoUsers(ctx context.Context, id string) error
	// Update
	UpdateSignups(ctx context.Context, signup *models.Signup) error
	UpdateLogins(ctx context.Context, login *models.Login) error
	UpdateForms(ctx context.Context, form *models.Form) error
	UpdateDashboards(ctx context.Context, dashboard *models.Dashboard) error
	UpdateAuths(ctx context.Context, auth *models.Auth) error
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

/* Gets */

// GetKeyValueById do which the implementation
func GetKeyValueById(ctx context.Context, id string) (*models.KeyValue, error) {
	return implementation.GetKeyValueById(ctx, id)
}

// GetFormByTitle do which the implementation
func GetFormByTitle(ctx context.Context, title string) (*models.Form, error) {
	return implementation.GetFormByTitle(ctx, title)
}

// GetFormById do wich the implementation
func GetFormById(ctx context.Context, id string) (*models.Form, error) {
	return implementation.GetFormById(ctx, id)
}

// GetLoginByAuthId do wich the implementation
func GetLoginByAuthId(ctx context.Context, auth_id string) (*models.Login, error) {
	return implementation.GetLoginByAuthId(ctx, auth_id)
}

// GetSignupById do wich the implementation
func GetSignupById(ctx context.Context, id string) (*models.Signup, error) {
	return implementation.GetSignupById(ctx, id)
}

// GetSignupByFormId do wich the implementation
func GetSignupByFormId(ctx context.Context, form_id string) (*models.Signup, error) {
	return implementation.GetSignupByFormId(ctx, form_id)
}

// GetDashboardByAuthId do wich the implementation
func GetDashboardByAuthId(ctx context.Context, auth_id string) (*models.Dashboard, error) {
	return implementation.GetDashboardByAuthId(ctx, auth_id)
}

// GetAuthById do wich the implementation
func GetAuthById(ctx context.Context, id string) (*models.Auth, error) {
	return implementation.GetAuthById(ctx, id)
}

// GetAuthByEmail do wich the implementation
func GetAuthByEmail(ctx context.Context, email string) (*models.Auth, error) {
	return implementation.GetAuthByEmail(ctx, email)
}

// GetUserByNickName do wich the implementation
func GetUserByNickName(ctx context.Context, nick_name string) (*models.User, error) {
	return implementation.GetUserByNickName(ctx, nick_name)
}

// GetUserById do wich the implementation
func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

// GetInfoUserByUserId do which the implementation
func GetInfoUserByUserId(ctx context.Context, user_id string) (*models.InfoUser, error) {
	return implementation.GetInfoUserByUserId(ctx, user_id)
}

// GetInfoUserByPhone do which the implementation
func GetInfoUserByPhone(ctx context.Context, phone int) (*models.InfoUser, error) {
	return implementation.GetInfoUserByPhone(ctx, phone)
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

// GetQueryValueByUserId do which the implementation
func GetQueryValueByUserId(ctx context.Context, user_id string) ([]*models.QueryValue, error) {
	return implementation.GetQueryValueByUserId(ctx, user_id)
}

// GetEndPointByUrlId do which the implementation
func GetEndPointByUrlId(ctx context.Context, url_id string) ([]*models.Endpoint, error) {
	return implementation.GetEndPointByUrlId(ctx, url_id)
}

/* Lists */

// ListForms do which the implementation
func ListForms(ctx context.Context) ([]*models.Form, error) {
	return implementation.ListForms(ctx)
}

// ListLogins do which the implementation
func ListLogins(ctx context.Context) ([]*models.Login, error) {
	return implementation.ListLogins(ctx)
}

// ListSignups do which the implementation
func ListSignups(ctx context.Context) ([]*models.Signup, error) {
	return implementation.ListSignups(ctx)
}

// ListDashboards do which the implementation
func ListDashboards(ctx context.Context) ([]*models.Dashboard, error) {
	return implementation.ListDashboards(ctx)
}

// ListAuths do which the implementation
func ListAuths(ctx context.Context) ([]*models.Auth, error) {
	return implementation.ListAuths(ctx)
}

// ListProtocols do which the implementation
func ListProtocols(ctx context.Context) ([]*models.Protocol, error) {
	return implementation.ListProtocols(ctx)
}

// ListURLs do which the implementation
func ListURLs(ctx context.Context) ([]*models.Url, error) {
	return implementation.ListURLs(ctx)
}

// ListQueryKeys do which the implementation
func ListQueryKeys(ctx context.Context) ([]*models.QueryKey, error) {
	return implementation.ListQueryKeys(ctx)
}

// ListQueryValues do which the implementation
func ListQueryValues(ctx context.Context) ([]*models.QueryValue, error) {
	return implementation.ListQueryValues(ctx)
}

// ListEndpoints do which the implementation
func ListEndpoints(ctx context.Context) ([]*models.Endpoint, error) {
	return implementation.ListEndpoints(ctx)
}

// ListUsers do which the implementation
func ListUsers(ctx context.Context) ([]*models.User, error) {
	return implementation.ListUsers(ctx)
}

// ListInfoUsers do which the implementation
func ListInfoUsers(ctx context.Context) ([]*models.InfoUser, error) {
	return implementation.ListInfoUsers(ctx)
}

/* Inserts */

// InsertIntoSignups do which the implementation
func InsertIntoSignups(ctx context.Context, signup *models.Signup) error {
	return implementation.InsertIntoSignups(ctx, signup)
}

// InsertIntoLogins do which the implementation
func InsertIntoLogins(ctx context.Context, login *models.Login) error {
	return implementation.InsertIntoLogins(ctx, login)
}

// InsertIntoForms do which the implementation
func InsertIntoForms(ctx context.Context, form *models.Form) error {
	return implementation.InsertIntoForms(ctx, form)
}

// InsertIntoDashboards do which the implementation
func InsertIntoDashboards(ctx context.Context, dashboard *models.Dashboard) error {
	return implementation.InsertIntoDashboards(ctx, dashboard)
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

/* Deletes */

// DeleteLogins do which the implementation
func DeleteLogins(ctx context.Context, id string) error {
	return implementation.DeleteLogins(ctx, id)
}

// DeleteSignups do which the implementation
func DeleteSignups(ctx context.Context, id string) error {
	return implementation.DeleteSignups(ctx, id)
}

// DeleteForms do which the implementation
func DeleteForms(ctx context.Context, id string) error {
	return implementation.DeleteForms(ctx, id)
}

// DeleteDashboards do which the implementation
func DeleteDashboards(ctx context.Context, id string) error {
	return implementation.DeleteDashboards(ctx, id)
}

// DeleteAuths do which the implementation
func DeleteAuths(ctx context.Context, id string) error {
	return implementation.DeleteAuths(ctx, id)
}

// DeleteProtocols do which the implementation
func DeleteProtocols(ctx context.Context, id string) error {
	return implementation.DeleteProtocols(ctx, id)
}

// DeleteURLs do which the implementation
func DeleteURLs(ctx context.Context, id string) error {
	return implementation.DeleteURLs(ctx, id)
}

// DeleteEndpoints do which the implementation
func DeleteEndpoints(ctx context.Context, id string) error {
	return implementation.DeleteEndpoints(ctx, id)
}

// DeleteQueryKeys do which the implementation
func DeleteQueryKeys(ctx context.Context, id string) error {
	return implementation.DeleteQueryKeys(ctx, id)
}

// DeleteQueryValues do which the implementation
func DeleteQueryValues(ctx context.Context, id string) error {
	return implementation.DeleteQueryValues(ctx, id)
}

// DeleteUsers do which the implementation
func DeleteUsers(ctx context.Context, id string) error {
	return implementation.DeleteUsers(ctx, id)
}

// DeleteInfoUsers do which the implementation
func DeleteInfoUsers(ctx context.Context, id string) error {
	return implementation.DeleteInfoUsers(ctx, id)
}

/* Updates */

// UpdateLogins do which the implementation
func UpdateLogins(ctx context.Context, login *models.Login) error {
	return implementation.UpdateLogins(ctx, login)
}

// UpdateSignups do which the implementation
func UpdateSignups(ctx context.Context, signup *models.Signup) error {
	return implementation.UpdateSignups(ctx, signup)
}

// UpdateForms do which the implementation
func UpdateForms(ctx context.Context, form *models.Form) error {
	return implementation.UpdateForms(ctx, form)
}

// UpdateDashboards do which the implementation
func UpdateDashboards(ctx context.Context, dashboard *models.Dashboard) error {
	return implementation.UpdateDashboards(ctx, dashboard)
}

// UpdateAuths do which the implementation
func UpdateAuths(ctx context.Context, auth *models.Auth) error {
	return implementation.UpdateAuths(ctx, auth)
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
