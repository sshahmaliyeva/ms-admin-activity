package  main

import(
	log "github.com/sirupsen/logrus"
	"github.com/sshahmaliyeva/ms-admin-activity/tree/master/config"
	"github.com/joho/godotenv"
	"net/http"

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

	// TODO: Add Http Listener to port 80
}

/*	TODO LIST:
*	- (Optional) Install any todo plugin or highlighter to view list of todos in project
*	- Add go.mod *DONE*
*	- Add README file
*	- Add .gitignore for go project (can copy from ms-sign-settings)
*	- Add Docker file (can copy from ms-sign-settings but needs refactoring)
*	- Add infra terraform
*	- Add circle CI (copy but refactor)
*/
