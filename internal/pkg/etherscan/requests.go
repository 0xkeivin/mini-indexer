package etherscan

import (
	"io"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// Polls etherscan for new transactions records
func SendGetReq(url string) ([]byte, error) {
	// Create a new HTTP client and set a timeout
	client := &http.Client{
		Timeout: time.Second * 10, // Set Timeout to 5 seconds
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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Info("Error reading response body: %s", err)
	}
	// Return the response body
	return body, nil

}
