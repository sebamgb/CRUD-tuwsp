package models

import "github.com/golang-sql/civil"

type Dashboard struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Menu  string `json:"menu"`
	App   string `json:"app"`
	Owner string `json:"owner"`
}

type Signup struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	NickName        string `json:"nick_name"`
	Email           string `json:"email"`
	Phone           int    `json:"phone"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	FormId          string `json:"form_id"`
}

type Login struct {
	Id        string         `json:"id"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt civil.DateTime `json:"created_at"`
	LogOut    civil.DateTime `json:"log_out"`
	AuthId    string         `json:"auth_id"`
	FormId    string         `json:"form_id"`
}

type Form struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	App       string `json:"app"`
	Key_Value string `json:"key_value"`
}
