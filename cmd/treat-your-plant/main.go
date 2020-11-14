package main

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{PrettyPrint: true})
	log.SetLevel(log.DebugLevel)

	log.Info("Hello, i have nothing to show you yet :(")

	dbUrl := "We don't have one yet..."

	db, err := sqlx.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(10)

}
