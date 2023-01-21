package models

type Url struct {
	Id         string `url:"id"`
	Domain     string `url:"domain"`
	ProtocolId string `url:"protocol_id"`
}

type Protocol struct {
	Id       string `url:"id"`
	Protocol string `url:"protocol"`
}

type QueryKey struct {
	Id       int    `url:"id"`
	KeyParam string `url:"key_param"`
	UrlId    string `url:"url_id"`
}

type QueryValue struct {
	Id         int    `url:"id"`
	ValueParam string `url:"value_param"`
	UrlId      string `url:"url_id"`
}

type Endpoint struct {
	Id       int    `url:"id"`
	Endpoint string `url:"endpoint"`
	UrlId    string `url:"url_id"`
}
