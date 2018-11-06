package sorter

import (
	"testing"

	"geo-fence-siever/src/loader"

	"github.com/stretchr/testify/assert"
)

func TestCustomerSortedSlice_InsertAndString(t *testing.T) {
	testCases := []struct {
		customer            *loader.Customer
		customerSortedSlice string
	}{
		{
			customer: &loader.Customer{
				UserID: 23,
				Name:   "Christina McArdle",
			},
			customerSortedSlice: "User ID: 23, Name: Christina McArdle\n",
		},
		{
			customer: &loader.Customer{
				UserID: 11,
				Name:   "Alice Cahill",
			},
			customerSortedSlice: "User ID: 11, Name: Alice Cahill\nUser ID: 23, Name: Christina McArdle\n",
		},
		{
			customer: &loader.Customer{
				UserID: 37,
				Name:   "Jack Enright",
			},
			customerSortedSlice: "User ID: 11, Name: Alice Cahill\nUser ID: 23, Name: Christina McArdle\nUser ID: 37, Name: Jack Enright\n",
		},
		{
			customer: &loader.Customer{
				UserID: 15,
				Name:   "Ian McArdle",
			},
			customerSortedSlice: "User ID: 11, Name: Alice Cahill\nUser ID: 15, Name: Ian McArdle\nUser ID: 23, Name: Christina McArdle\nUser ID: 37, Name: Jack Enright\n",
		},
	}

	customerSortedSlice := NewCustomerSortedSlice()

	for _, testCase := range testCases {
		customerSortedSlice.Insert(testCase.customer)
		assert.Equal(t, testCase.customerSortedSlice, customerSortedSlice.String())
	}
}
