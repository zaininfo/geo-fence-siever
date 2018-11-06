package main

import (
	"flag"
	"os"

	"geo-fence-siever/src/application"
)

var app *application.Application

func init() {
	app = createApplication()
}

func createApplication() *application.Application {
	filename := flag.String("filename", "", "The text file that contains the customer data. (Required)")
	flag.Parse()

	if *filename == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	return application.NewApplication(&application.Config{
		DataFilename:    *filename,
	})
}

func main() {
	app.Run()
}
