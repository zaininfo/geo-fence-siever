package application

import (
	"fmt"
	"log"

	"geo-fence-siever/src/calculator"
	"geo-fence-siever/src/loader"
	"geo-fence-siever/src/sorter"
)

const officeLatitude = 52.5200
const officeLongitude = 13.4050
const luckyRadius = 100

// Config contains all the configurations of application
type Config struct {
	DataFilename string
}

// Application contains values required for running application
type Application struct {
	dataFilename string
	loader       loader.DataLoader
	calculator   calculator.Calculator
	sortedSlice  sorter.SortedSlice
}

// NewApplication creates and returns application
func NewApplication(c *Config) *Application {
	customerLoader := loader.NewCustomerLoader(loader.CustomerLoaderConfig{
		DataFilename: c.DataFilename,
	})

	return &Application{
		dataFilename: c.DataFilename,
		loader:       customerLoader,
		calculator:   calculator.NewGreatCircleCalculator(),
		sortedSlice:  sorter.NewCustomerSortedSlice(),
	}
}

// Run starts the processing
func (a *Application) Run() {
	customers, err := a.loader.Load()
	if err != nil {
		log.Fatal(err)
	}

	coordinate1 := calculator.Coordinates{
		Latitude:  officeLatitude,
		Longitude: officeLongitude,
	}

	for _, customer := range customers {
		customerStructure := customer.(*loader.Customer)

		coordinate2 := calculator.Coordinates{
			Latitude:  customerStructure.Latitude,
			Longitude: customerStructure.Longitude,
		}

		greatCircleDistance := a.calculator.Calculate(coordinate1, coordinate2)

		if greatCircleDistance <= luckyRadius {
			a.sortedSlice.Insert(customerStructure)
		}
	}

	fmt.Print(a.sortedSlice.String())
}
