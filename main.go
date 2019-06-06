package  main

import(
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv" 
	"github.com/go-chi/chi/middleware"
	"github.com/go-pg/pg"
)

var db *pg.DB

func init() {
	godotenv.Load()
}

func main(){
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Application starting...")

	log.Info("Database connection verified...")

	router := mux.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	router.HandleFunc("/readiness", health)
	router.HandleFunc("/health", health)

	log.Info("Server listen at :80")
	log.Fatal(http.ListenAndServe(":80", router))
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK!")
}
