package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/0xkeivin/web3-indexer/internal/pkg/db"
	"github.com/0xkeivin/web3-indexer/internal/pkg/env"
	"github.com/0xkeivin/web3-indexer/internal/pkg/etherscan"
	"github.com/0xkeivin/web3-indexer/internal/utils"

	log "github.com/sirupsen/logrus"
)

// create struct for json response
type Response struct {
	Status  string             `json:"status"`
	Message string             `json:"message"`
	Result  []db.BlockChainLog `json:"result"`
}

func main() {
	// Initialize env variables
	loadedConfig, err := env.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}
	log.Info("Initialized env variables...")

	// log.Infof("ETHERSCAN_API_KEY: %v", loadedConfig.ETHERSCAN_API_KEY)
	log.Infof("DBName: %v", loadedConfig.DBName)
	// Start web3-indexer
	log.Info("Starting web3-indexer...")
	// Using USDC contract address
	// address := "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
	// Base URL
	baseUrl := `https://api.etherscan.io/api?module=logs&action=getLogs&page=1&offset=1000`
	// Initialize database
	db.ConnectDB(&loadedConfig)
	db.AutoMigrate(db.DB)
	// create slice of objects
	var response Response
	// for loop
	ticker := time.Tick(10 * time.Second)
	for range ticker {
		currentTime := time.Now()
		// print log with tick time
		log.Infof("Tick at - %s", currentTime.Format("2006-01-02 15:04:05"))
		// get latest block number
		latestBlockNumber := etherscan.GetLatestBlockNumber(loadedConfig.ETHERSCAN_API_KEY)
		fromBlock := latestBlockNumber - 10 // get last 10 blocks
		log.Infof("Latest block number: %v", latestBlockNumber)
		log.Infof("From block number: %v", fromBlock)

		fullUrl := baseUrl + `&address=` + loadedConfig.CONTRACT_ADDRESS + `&fromBlock=` + fmt.Sprint(fromBlock) + `&toBlock=` + fmt.Sprint(latestBlockNumber) + `&apikey=` + loadedConfig.ETHERSCAN_API_KEY
		log.Infof("Full URL: %s", fullUrl)
		// Poll etherscan
		resp, err := etherscan.SendGetReq(fullUrl)
		if err != nil {
			log.Infof("Error polling etherscan: %s", err)
		}

		// Convert response to JSON
		utils.ConvertByteToJSON(resp)
		// unmarshal JSON to struct
		err = json.Unmarshal(resp, &response)
		if err != nil {
			log.Infof("Error unmarshalling JSON: %s", err)
		}
		// print response
		log.Infof("Response: %v", response.Message)
		// create slice of objects
		db.InsertObj(db.DB, response.Result)
	}
}
