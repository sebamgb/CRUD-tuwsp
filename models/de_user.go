package models

type InfoUser struct {
	Id         string `json:"id"`
	Phone      int    `json:"phone"`
	Country    string `json:"country"`
	CodCountry string `json:"cod_country"`
	Birthday   string `json:"birthday"`
	UserId     string `json:"user_id"`
}

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	NickName string `json:"nick_name"`
	UrlId    string `json:"url_id"`
}
