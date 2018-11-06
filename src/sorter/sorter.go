package sorter

import (
	"fmt"
	"sort"

	"geo-fence-siever/src/loader"
)

// SortedSlice defines the interface for storing sorted data
type SortedSlice interface {
	Insert(interface{})
	String() string
}

type customerSortedSlice struct {
	sortedSlice []*loader.Customer
}

// NewCustomerSortedSlice creates and returns a data structure for storing sorted customer data
func NewCustomerSortedSlice() SortedSlice {
	return &customerSortedSlice{
		sortedSlice: []*loader.Customer{},
	}
}

// Insert adds a new customer to the storage according to its user ID
func (s *customerSortedSlice) Insert(customer interface{}) {
	customerStructure := customer.(*loader.Customer)
	length := len(s.sortedSlice)

	fitsAt := sort.Search(length, func(index int) bool {
		return s.sortedSlice[index].UserID >= customerStructure.UserID
	})

	newSortedSlice := make([]*loader.Customer, fitsAt+1)

	copy(newSortedSlice, s.sortedSlice[:fitsAt])
	newSortedSlice[fitsAt] = customerStructure
	newSortedSlice = append(newSortedSlice, s.sortedSlice[fitsAt:]...)

	s.sortedSlice = newSortedSlice
}

// String returns the data in the storage encoded as string
func (s *customerSortedSlice) String() string {
	var stringifiedCustomers string

	for _, customer := range s.sortedSlice {
		stringifiedCustomers += fmt.Sprintf("User ID: %d, Name: %s\n", customer.UserID, customer.Name)
	}

	return stringifiedCustomers
}
