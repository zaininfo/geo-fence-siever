package tests

import (
	"bytes"
	"io"
	"log"
	"os"
	"testing"

	"geo-fence-siever/src/application"

	"github.com/stretchr/testify/assert"
)

func TestEndToEnd_NoCustomers(t *testing.T) {
	app := application.NewApplication(&application.Config{
		DataFilename: "data1_test.txt",
	})

	output := captureOutput(app.Run)
	expectedOutput := ""

	assert.Equal(t, expectedOutput, output)
}

func TestEndToEnd_OnlyDistantCustomers(t *testing.T) {
	app := application.NewApplication(&application.Config{
		DataFilename: "data2_test.txt",
	})

	output := captureOutput(app.Run)
	expectedOutput := ""

	assert.Equal(t, expectedOutput, output)
}

func TestEndToEnd_BothDistantAndCloseCustomers(t *testing.T) {
	app := application.NewApplication(&application.Config{
		DataFilename: "data3_test.txt",
	})

	output := captureOutput(app.Run)
	expectedOutput := "User ID: 4, Name: Ian Kehoe\nUser ID: 12, Name: Christina McArdle\nUser ID: 17, Name: Patricia Cahill\n"

	assert.Equal(t, expectedOutput, output)
}

// Adapted from: https://stackoverflow.com/a/10476304
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout = w

	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	f()

	w.Close()
	os.Stdout = old
	out := <-outC

	return out
}
