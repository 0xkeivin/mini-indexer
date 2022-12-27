package main

import (
	"fmt"
	"time"

	"github.com/0xkeivin/web3-indexer/internal/pkg/db"
	"github.com/0xkeivin/web3-indexer/internal/pkg/env"
	"github.com/0xkeivin/web3-indexer/internal/pkg/etherscan"
	"github.com/0xkeivin/web3-indexer/internal/utils"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Starting web3-indexer...")
	etherscanAPIKey := env.GoDotEnvVariable("ETHERSCAN_API_KEY")

	// log.Infof("Using Etherscan API Key: %s", etherscanAPIKey)

	// Using USDC contract address
	address := "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
	// Base URL
	baseUrl := `https://api.etherscan.io/api?module=logs&action=getLogs&page=1&offset=1000`

	// test section
	db, err := db.ConnectToDB()
	if err != nil {
		log.Infof("Error connecting to database: %s", err)
	}
	defer db.Close()
	log.Infof("DB: %v", db)

	// for loop
	ticker := time.Tick(20 * time.Second)
	for range ticker {
		currentTime := time.Now()
		// print log with tick time
		log.Infof("Tick at - %s", currentTime.Format("2006-01-02 15:04:05"))
		// get latest block number
		latestBlockNumber := etherscan.GetLatestBlockNumber(etherscanAPIKey)
		fromBlock := latestBlockNumber - 10 // get last 10 blocks
		log.Infof("Latest block number: %v", latestBlockNumber)
		log.Infof("From block number: %v", fromBlock)

		// fullUrl := baseUrl + `&address=` + address + `&apikey=` + etherscanAPIKey
		fullUrl := baseUrl + `&address=` + address + `&fromBlock=` + fmt.Sprint(fromBlock) + `&toBlock=` + fmt.Sprint(latestBlockNumber) + `&apikey=` + etherscanAPIKey
		log.Infof("Full URL: %s", fullUrl)
		// Poll etherscan
		resp, err := etherscan.SendGetReq(fullUrl)
		if err != nil {
			log.Infof("Error polling etherscan: %s", err)
		}
		// Convert response to JSON
		utils.ConvertByteToJSON(resp)

	}
}
