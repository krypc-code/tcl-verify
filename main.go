package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"log"
)

func main() {
	// read data from yaml file
	data := ReadYaml("data.yaml")

	payloadBytes, err := json.Marshal(data.Payload)
	if err != nil {
		log.Panic("Error while marshalling data to json")
	}
	// get hmac hash of payloadBytes using hmac key
	dataHash := GetHmacHash(payloadBytes, data.HMAC)
	log.Println("data hash : ", dataHash)
	if dataHash != data.DataHash {
		log.Println("Data hash does not match")
	} else {
		log.Println("Data hash matches")
	}
}

// GetHmacHash use hmac key to generate payload hash
func GetHmacHash(payload []byte, hmacKey string) string {
	hmac := hmac.New(sha256.New, []byte(hmacKey))

	_, err := hmac.Write(payload)
	if err != nil {
		log.Panic("Error while hashing data")
	}
	sum := hmac.Sum(nil)
	return base64.StdEncoding.EncodeToString(sum)
}
