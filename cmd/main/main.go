package main

import (
	"fmt"
	"time"

	"github.com/0xkeivin/web3-indexer/internal/pkg/env"
	"github.com/0xkeivin/web3-indexer/internal/pkg/etherscan"
	"github.com/0xkeivin/web3-indexer/internal/utils"

	log "github.com/sirupsen/logrus"
)

// // Create a function that converts byte[] to JSON
// func convertByteToJSON(body []byte) {
// 	var data interface{}
// 	err := json.Unmarshal(body, &data)
// 	if err != nil {
// 		log.Infof("Error converting byte to JSON: %s", err)
// 		return
// 	}
// 	// Encode data as JSON
// 	jsonData, err := json.MarshalIndent(data, "", " ")
// 	if err != nil {
// 		log.Infof("Error encoding data as JSON: %s", err)
// 		return
// 	}
// 	// Write JSON to file
// 	err = os.WriteFile("response.json", jsonData, 0644)
// 	if err != nil {
// 		log.Infof("Error writing JSON to file: %s", err)
// 		return
// 	}
// }

// use godot package to load/read the .env file and
// return the value of the key
// func goDotEnvVariable(key string) string {

// 	// load .env file
// 	err := godotenv.Load(".env")

// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}

// 	return os.Getenv(key)
// }

// // Polls etherscan for new transactions records
// func pollEtherscan(url string) ([]byte, error) {
// 	// Create a new HTTP client and set a timeout
// 	client := &http.Client{
// 		Timeout: time.Second * 10, // Set Timeout to 5 seconds
// 	}
// 	// Create a new request
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		// return "", err
// 		log.Info("Error creating request: %s", err)
// 	}
// 	// Set headers, if necessary
// 	req.Header.Set("Content-Type", "application/json")
// 	// Send the request
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Info("Error sending request: %s", err)
// 	}
// 	// Read the response body
// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Info("Error reading response body: %s", err)
// 	}
// 	// Return the response body
// 	return body, nil

// }

// type BlockResponse struct {
// 	Status  string `json:"status"`
// 	Message string `json:"message"`
// 	Result  string `json:"result"`
// }

// // function to get latest block number
// func getLatestBlockNumber(apiKey string) int {
// 	// get current time
// 	currentTime := time.Now().Unix()
// 	log.Infof("Current time: %d", currentTime)
// 	// get latest block number
// 	baseUrl := `https://api.etherscan.io/api?module=block&action=getblocknobytime&timestamp=`
// 	fullUrl := baseUrl + fmt.Sprint(currentTime) + `&closest=before&apikey=` + apiKey
// 	log.Infof("Full URL: %s", fullUrl)
// 	// Poll etherscan
// 	resp, err := pollEtherscan(fullUrl)
// 	if err != nil {
// 		log.Infof("Error polling etherscan: %s", err)
// 	}
// 	// Convert response to JSON
// 	var blockResponse BlockResponse
// 	err = json.Unmarshal(resp, &blockResponse)
// 	if err != nil {
// 		log.Infof("Error converting response to JSON: %s", err)
// 	}
// 	if blockResponse.Message != "OK" {
// 		log.Infof("Error getting latest block number: %s", blockResponse.Message)
// 	}
// 	// Return latest block number
// 	latestBlock, err := strconv.Atoi(blockResponse.Result)
// 	if err != nil {
// 		log.Infof("Error converting block number to int: %s", err)
// 	}
// 	return latestBlock

// }
func main() {
	log.Info("Starting web3-indexer...")
	// etherscanAPIKey := goDotEnvVariable("ETHERSCAN_API_KEY")
	etherscanAPIKey := env.GoDotEnvVariable("ETHERSCAN_API_KEY")

	// log.Infof("Using Etherscan API Key: %s", etherscanAPIKey)

	// Using USDC contract address
	address := "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
	// Base URL
	baseUrl := `https://api.etherscan.io/api?module=logs&action=getLogs&page=1&offset=1000`

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
		// convertByteToJSON(resp)

	}
}
