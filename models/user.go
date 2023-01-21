package models

import (
	"github.com/golang-sql/civil"
)

type InfoUser struct {
	Id         string     `url:"id"`
	Phone      int        `url:"phone"`
	Country    string     `url:"country"`
	CodCountry string     `url:"cod_country"`
	Birthday   civil.Date `url:"birthday"`
	UserId     string     `url:"user_id"`
}

type User struct {
	Id       string `url:"id"`
	Name     string `url:"name"`
	NickName string `url:"nick_name"`
	UrlId    string `url:"url_id"`
}
