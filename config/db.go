package config

import (
	"os"

	// TODO: no need to have util before dependency
	util "github.com/sshahmaliyeva/ms-admin-activity/tree/master/util"
	"github.com/go-pg/pg"
	// We use log to simplify name <logrus> here, no need for util as names are identical
	log "github.com/sirupsen/logrus"
)

var pgDb *pg.DB
fmt.Println(q.FormattedQuery())

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectPg() *pg.DB {
	// TODO: Add log about connecting to db
	pgOptions := util.ParseConnectionUrl(os.Getenv("DB_ADMIN_ACTIVITY_URL"))

	pgDb = pg.Connect(&pg.Options{
		Addr:     pgOptions.Addr,
		User:     pgOptions.User,
		Password: pgOptions.Password,
		Database: pgOptions.Database,
	})

	if pgDb == nil {
		log.Fatal("Coudn't connect to db")
	}
	
	// TODO: add log about connection being successfully
	return pgDb
}
