package models

import "github.com/golang-sql/civil"

type KeyValue struct {
	Id                         string `json:"id"`
	Title                      string `json:"title"` // Form title
	LabelName                  string `json:"label_name"`
	PlaceholderName            string `json:"placeholder_name"`
	LabelNickname              string `json:"label_nickname"`
	PlaceholderNickname        string `json:"placeholder_nickname"`
	LabelEmail                 string `json:"label_email"`
	PlaceholderEmail           string `json:"placeholder_email"`
	LabelPhone                 string `json:"label_phone"`
	PlaceholderPhone           string `json:"placeholder_phone"`
	LabelBirthday              string `json:"label_birthday"`
	LabelCountry               string `json:"label_country"`
	PlaceholderCountry         string `json:"placeholder_country"`
	LabelPassword              string `json:"label_password"`
	PlaceholderPassword        string `json:"placeholder_password"`
	LabelConfirmPassword       string `json:"label_confirm_password"`
	PlaceholderConfirmPassword string `json:"placeholder_confirm_password"`
	InputSubmit                string `json:"input_submit"`
	LabelId                    string `json:"label_id"`
	Author                     string `json:"author"`
}

type Dashboard struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Menu  string `json:"key_value"`
	App   string `json:"app"`
	Owner string `json:"owner"`
}

type Signup struct {
	Id              string `json:"id"`
	Title           string `json:"title"`
	Url             string `json:"url"`
	Method          string `json:"method"`
	Name            string `json:"name"`
	NickName        string `json:"nick_name"`
	Email           string `json:"email"`
	Phone           int    `json:"phone"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Country         string `json:"country"`
	FormId          string `json:"form_id"`
}

type Login struct {
	Id        string         `json:"id"`
	Title     string         `json:"title"`
	Url       string         `json:"url"`
	Method    string         `json:"method"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt civil.DateTime `json:"created_at"`
	LogOut    civil.DateTime `json:"log_out"`
	AuthId    string         `json:"auth_id"`
	FormId    string         `json:"form_id"`
}

type Form struct {
	Id        string `json:"id"`
	App       string `json:"app"`
	Key_Value string `json:"key_value"`
	Author    string `json:"author"`
}
