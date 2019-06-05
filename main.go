package  main

import(
	log "github.com/sirupsen/logrus"
	"github.com/sshahmaliyeva/ms-admin-activity/blob/master/config"
	"github.com/joho/godotenv"

	"github.com/go-pg/pg"
)

var db *pg.DB

func init() {
	godotenv.Load()
}

func main(){
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Application starting...")

	db = config.ConnectPg()
	log.Info("Database connection verified...")
}
