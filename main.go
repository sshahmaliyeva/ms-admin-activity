package  main

import(
	"fmt"
	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main(){
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Application starting...")
}
