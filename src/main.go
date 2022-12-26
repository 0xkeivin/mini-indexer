package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// Create a function that converts byte[] to JSON
func convertByteToJSON(body []byte) {
	var data interface{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Infof("Error converting byte to JSON: %s", err)
		return
	}
	// Encode data as JSON
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Infof("Error encoding data as JSON: %s", err)
		return
	}
	// Write JSON to file
	err = ioutil.WriteFile("response.json", jsonData, 0644)
	if err != nil {
		log.Infof("Error writing JSON to file: %s", err)
		return
	}
}

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// Polls etherscan for new transactions records
func pollEtherscan(url string) ([]byte, error) {
	// Create a new HTTP client and set a timeout
	client := &http.Client{
		Timeout: time.Second * 5, // Set Timeout to 5 seconds
	}
	// Create a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// return "", err
		log.Info("Error creating request: %s", err)
	}
	// Set headers, if necessary
	req.Header.Set("Content-Type", "application/json")
	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Info("Error sending request: %s", err)
	}
	// Read the response body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Info("Error reading response body: %s", err)
	}
	// Return the response body
	return body, nil

}
func main() {
	etherscanAPIKey := goDotEnvVariable("ETHERSCAN_API_KEY")
	// log.Infof("Using Etherscan API Key: %s", etherscanAPIKey)
	// Using USDC contract address
	address := "0x7EA2be2df7BA6E54B1A9C70676f668455E329d29"
	// Base URL
	baseUrl := `https://api.etherscan.io/api?module=logs&action=getLogs&page=1&offset=1000`
	fullUrl := baseUrl + `&address=` + address + `&apikey=` + etherscanAPIKey
	log.Infof("Full URL: %s", fullUrl)
	// Poll etherscan
	resp, err := pollEtherscan(fullUrl)
	if err != nil {
		log.Info("Error polling etherscan: %s", err)
	}
	// Convert response to JSON
	convertByteToJSON(resp)
}
