package controllers

import (
	"fmt"
	"net/http"
	"01cloud-payment/internal/models"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	fmt.Println(DBURL)
	server.DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		log.Fatal("Cannot connect to %s database: %v", Dbdriver, err)
	}

	server.Migrate()

	server.Router = mux.NewRouter()

	server.initializeRoutes()

}
func (server *Server) Migrate() {
	server.DB.AutoMigrate(
		models.User{},
		models.Project{},
		models.Subscription{},
	)
}

func (server *Server) Run(addr string) {

	log.Info("Listening to port http://localhost:8081")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
