package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"log"
	"tcl-verify/config"
)

func main() {
	// read data from yaml file
	data := config.ReadYaml("data.yaml")

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
	keyRaw, err := base64.StdEncoding.DecodeString(hmacKey)
	if err != nil {
		log.Panic("Error while decoding hmac key, " + err.Error())
	}
	hmac := hmac.New(sha256.New, keyRaw)

	_, err = hmac.Write(payload)
	if err != nil {
		log.Panic("Error while hashing data")
	}
	sum := hmac.Sum(nil)
	return base64.StdEncoding.EncodeToString(sum)
}
