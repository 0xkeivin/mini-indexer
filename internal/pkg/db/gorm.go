package db

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/0xkeivin/web3-indexer/internal/pkg/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB connects to the database
func ConnectDB(config *env.Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.WithFields(log.Fields{
			"DB-INIT": "FAILED",
		}).Fatal("Failed to connect to database!")

	}
	log.WithFields(log.Fields{
		"DB-INIT": "OK",
	}).Info("Successfully connected to database!")
}

// AutoMigrate will migrate the schema for you
func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&BlockChainLog{})
	if err != nil {
		log.WithFields(log.Fields{
			"DB-AUTOMIGRATE": "FAILED",
		}).Fatal("Failed to migrate database!")
	}
	log.WithFields(log.Fields{
		"DB-AUTOMIGRATE": "OK",
	}).Info("DB migrated!")
}

// PopulateDB populates the database with the latest block number
// func PopulateDB (db *gorm.DB,  ) error {

// }
