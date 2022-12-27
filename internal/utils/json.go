package utils

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
)

// Create a function that converts byte[] to JSON
func ConvertByteToJSON(body []byte) {
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
	err = os.WriteFile("response.json", jsonData, 0644)
	if err != nil {
		log.Infof("Error writing JSON to file: %s", err)
		return
	}
}
