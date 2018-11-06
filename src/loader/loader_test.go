package loader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomerLoader_Load(t *testing.T) {
	customers := []*Customer{
		{
			UserID:    12,
			Name:      "Christina McArdle",
			Latitude:  52.986375,
			Longitude: -6.043701,
		},
		{
			UserID:    1,
			Name:      "Alice Cahill",
			Latitude:  51.92893,
			Longitude: -10.27699,
		},
		{
			UserID:    2,
			Name:      "Ian McArdle",
			Latitude:  51.8856167,
			Longitude: -10.4240951,
		},
	}

	customerLoader := NewCustomerLoader(CustomerLoaderConfig{
		DataFilename: "data_test.txt",
	})

	data, err := customerLoader.Load()
	if err != nil {
		t.Fatal("Data loading failed: ", err)
	}

	for index, item := range data {
		assert.Equal(t, customers[index], item.(*Customer))
	}
}
