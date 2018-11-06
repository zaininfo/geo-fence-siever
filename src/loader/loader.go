package loader

import (
	"bufio"
	"encoding/json"
	"os"
)

// DataLoader defines the interface for loading data
type DataLoader interface {
	Load() ([]interface{}, error)
}

// Customer defines the structure of customer data
type Customer struct {
	UserID    int64   `json:"user_id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude,string"`
	Longitude float64 `json:"longitude,string"`
}

type customerLoader struct {
	dataFilename string
}

// CustomerLoaderConfig defines the structure of configurations for customer loader
type CustomerLoaderConfig struct {
	DataFilename string
}

// NewCustomerLoader creates and returns a loader for customer data
func NewCustomerLoader(c CustomerLoaderConfig) DataLoader {
	return &customerLoader{
		dataFilename: c.DataFilename,
	}
}

// Load reads and returns all customers
func (l *customerLoader) Load() ([]interface{}, error) {
	file, err := os.Open(l.dataFilename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var customers []interface{}

	for scanner.Scan() {
		customer := &Customer{}

		err = json.Unmarshal(scanner.Bytes(), customer)
		if err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}
