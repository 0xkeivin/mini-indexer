package etherscan

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

// function to get latest block number
func GetLatestBlockNumber(apiKey string) int {
	// get current time
	currentTime := time.Now().Unix()
	log.Infof("Current time: %d", currentTime)
	// get latest block number
	baseUrl := `https://api.etherscan.io/api?module=block&action=getblocknobytime&timestamp=`
	fullUrl := baseUrl + fmt.Sprint(currentTime) + `&closest=before&apikey=` + apiKey
	log.Infof("Full URL: %s", fullUrl)
	// Poll etherscan
	resp, err := SendGetReq(fullUrl)
	if err != nil {
		log.Infof("Error polling etherscan: %s", err)
	}
	// Convert response to JSON
	var blockResponse BlockResponse
	err = json.Unmarshal(resp, &blockResponse)
	if err != nil {
		log.Infof("Error converting response to JSON: %s", err)
	}
	if blockResponse.Message != "OK" {
		log.Infof("Error getting latest block number: %s", blockResponse.Message)
	}
	// Return latest block number
	latestBlock, err := strconv.Atoi(blockResponse.Result)
	if err != nil {
		log.Infof("Error converting block number to int: %s", err)
	}
	return latestBlock

}
