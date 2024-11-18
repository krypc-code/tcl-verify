package config

import (
	"errors"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// PayloadData - represents a record whose hmac hash is pushed onto blockchain
type PayloadData struct {
	PeriodStartDate        string `json:"period_start_date" yaml:"period_start_date" `              // period start date
	PeriodEndDate          string `json:"period_end_date" yaml:"period_end_date"`                   // period end date
	InvoiceID              string `json:"invoice_id" yaml:"invoice_id"`                             // Invoice id as per invoice
	InvoiceAmount          string `json:"invoice_amount" yaml:"invoice_amount"`                     // Total invoice amount as per invoice
	TotalVolume            string `json:"total_volume,omitempty" yaml:"total_volume"`               // total minutes as per invoice
	PaymentInitiatedDate   string `json:"payment_initiated_date" yaml:"payment_initiated_date"`     // payment initiated date
	PaymentInitiatedAmount string `json:"payment_initiated_amount" yaml:"payment_initiated_amount"` // payment initiated amount
	ServiceType            string `json:"service_type" yaml:"service_type"`                         //VOICE / MMX /CHARGES
	TotalQuantity          string `json:"total_quantity,omitempty" yaml:"total_quantity"`           // Total SMS quantity
}

type VerifyData struct {
	Payload  PayloadData `json:"payload" yaml:"payload"`
	HMAC     string      `json:"hmac" yaml:"hmac"`
	DataHash string      `json:"data_hash" yaml:"data_hash"`
}

// ReadYaml - function for reading yaml
func ReadYaml(filePath string) VerifyData {

	var data VerifyData
	if strings.TrimSpace(filePath) == "" {
		log.Println("Empty value for CONFIG_FILE_PATH")
	}

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		log.Println("Yaml config file does not exist, will switch to environment variable configuration")
		return data

	}
	yamlByte := readConfigYamlFile(filePath)
	UnMarshalYaml(yamlByte, &data)
	return data
}

func readConfigYamlFile(filePath string) []byte {
	yamlByte, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("Error while reading config file ")
		log.Fatal(err)
	}
	return yamlByte
}

// UnMarshalYaml - function to unmarshal yaml
func UnMarshalYaml(yamlFileByte []byte, conf *VerifyData) {
	err1 := yaml.Unmarshal(yamlFileByte, &conf)
	if err1 != nil {
		log.Fatal("Error while unmarshalling bytes to yaml struct", err1.Error())
	}
}
