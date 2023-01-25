package models

type Url struct {
	Id         string `json:"id"`
	Domain     string `json:"domain"`
	ProtocolId string `json:"protocol_id"`
}

type Protocol struct {
	Id       string `json:"id"`
	Protocol string `json:"protocol"`
}

type QueryKey struct {
	Number   int    `json:"number"`
	Id       int    `json:"id"`
	KeyParam string `json:"key_param"`
	UrlId    string `json:"url_id"`
}

type QueryValue struct {
	Number     int    `json:"number"`
	Id         int    `json:"id"`
	ValueParam string `json:"value_param"`
	UserId     string `json:"user_id"`
}

type Endpoint struct {
	Number   int    `json:"number"`
	Id       int    `json:"id"`
	Endpoint string `json:"endpoint"`
	UrlId    string `json:"url_id"`
}
