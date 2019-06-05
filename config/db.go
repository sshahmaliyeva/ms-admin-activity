package config

import (
	"os"

	util "github.com/sshahmaliyeva/ms-admin-activity/util"
	"github.com/go-pg/pg"
	log "github.com/sirupsen/logrus"
)

var pgDb *pg.DB

// type dbLogger struct{}

// func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {}

// func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
// 	fmt.Println(q.FormattedQuery())
// }

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectPg() *pg.DB {
	log.Info(os.Getenv("DB_ADMIN_ACTIVITY_URL"))
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

	// pgDb.AddQueryHook(dbLogger{})

	return pgDb
}
