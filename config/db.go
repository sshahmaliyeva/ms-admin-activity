package config

import (
	// TODO: no need to have util before dependency
	//"github.com/sshahmaliyeva/ms-admin-activity/tree/master/util"
	"github.com/go-pg/pg"
	log "github.com/sirupsen/logrus"
)

var pgDb *pg.DB

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectPg() *pg.DB {
	log.Info("Database connection")
	//pgOptions := util.ParseConnectionUrl(os.Getenv("DB_ADMIN_ACTIVITY_URL"))

	pgDb = pg.Connect(&pg.Options{
		// Addr:     pgOptions.Addr,
		// User:     pgOptions.User,
		// Password: pgOptions.Password,
		// Database: pgOptions.Database,
	})

	if pgDb == nil {
		log.Fatal("Coudn't connect to db")
	}

	log.Info("Database connection successfully")
	return pgDb
}
