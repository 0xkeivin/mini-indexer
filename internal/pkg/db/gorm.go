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
func InsertObj(db *gorm.DB, objects []BlockChainLog) error {
	// create counter for number of objects inserted
	var objSkipped, objInserted int

	for _, obj := range objects {
		// using db.FirstOrCreate based on transaction_hash
		result := db.FirstOrCreate(&obj, BlockChainLog{TransactionHash: obj.TransactionHash})
		if result.Error != nil {
			log.WithFields(log.Fields{
				"DB-INSERT": "FAILED",
			}).Fatal("Failed to insert objects into DB!")
			return result.Error
		}
		if result.RowsAffected == 0 {
			objSkipped++
			log.WithFields(log.Fields{
				"DB-INSERT": "SKIPPED",
			}).Infof("Object already exists in DB! - %s", obj.TransactionHash)
			continue
		}
		objInserted++
		log.WithFields(log.Fields{
			"DB-INSERT": "OK",
		}).Infof("Inserted objects into DB!- %s", obj.TransactionHash)

	}
	log.WithFields(log.Fields{
		"DB-SKIP": objSkipped,
	}).Infof("Skipped object count into DB!- %v", objSkipped)
	log.WithFields(log.Fields{
		"DB-INSERT": objInserted,
	}).Infof("Inserted object count into DB!- %v", objInserted)
	return nil
}
