package util 

import (
	"strings"
	log "github.com/sirupsen/logrus"

)

func ParseConnectionUrl(origin_url string) DbOptions {
	var address = origin_url[strings.Index(origin_url, "@") + 1:strings.LastIndex(origin_url, "/")]
	log.Debug(address)

	var dbName = origin_url[strings.LastIndex(origin_url, "/") + 1:strings.LastIndex(origin_url, "?")]
	log.Debug(dbName)

	origin_url = strings.Replace(origin_url, "postgres://", "", 1)

	var user = origin_url[:strings.Index(origin_url, ":")]
	log.Debug(user)

	var password = origin_url[strings.Index(origin_url, ":") + 1:strings.Index(origin_url, "@")]
	log.Debug(password)

	return DbOptions{
		Addr: address,
		User: user,
		Password: password,
		Database: dbName,
	}
}

type DbOptions struct {
	Addr string
	User string
	Password string
	Database string
}