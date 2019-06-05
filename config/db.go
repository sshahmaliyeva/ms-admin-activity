package config

import (
	"os"

	util "github.com/sshahmaliyeva/ms-admin-activity/util"
	"github.com/go-pg/pg"
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

	return pgDb
}
