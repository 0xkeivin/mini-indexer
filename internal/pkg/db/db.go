package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func ConnectToDB() (*sql.DB, error) {
	db, err := sql.Open(
		"postgres",
		"host=localhost port=54320 user=user password=admin dbname=postgres sslmode=disable",
	)
	if err != nil {
		log.Errorf("Error connecting to database: %s", err)
		return nil, err
	}
	defer db.Close()
	// Test the connection to the database
	err = db.Ping()
	if err != nil {
		log.Infof("Error pinging database: %s", err)
		return nil, err
	}
	log.Info("Successfully connected to database!")
	return db, nil
}
