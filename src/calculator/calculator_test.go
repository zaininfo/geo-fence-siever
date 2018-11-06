package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreatCircleCalculator_Calculate(t *testing.T) {
	testCases := []struct {
		coordinate1         Coordinates
		coordinate2         Coordinates
		greatCircleDistance float64
	}{
		{
			coordinate1: Coordinates{
				Latitude:  0,
				Longitude: 0,
			},
			coordinate2: Coordinates{
				Latitude:  0,
				Longitude: 0,
			},
			greatCircleDistance: 0, // zero coordinates
		},
		{
			coordinate1: Coordinates{
				Latitude:  53.339428,
				Longitude: -6.257664,
			},
			coordinate2: Coordinates{
				Latitude:  53.339428,
				Longitude: -6.257664,
			},
			greatCircleDistance: 0, // same coordinates
		},
		{
			coordinate1: Coordinates{
				Latitude:  53.339428,
				Longitude: -6.257664,
			},
			coordinate2: Coordinates{
				Latitude:  -53.339450,
				Longitude: 173.742287,
			},
			greatCircleDistance: 20015.082725697488, // antipodal points
		},
		{
			coordinate1: Coordinates{
				Latitude:  53.339428,
				Longitude: -6.257664,
			},
			coordinate2: Coordinates{
				Latitude:  53.2451022,
				Longitude: -6.238335,
			},
			greatCircleDistance: 10.566936288867824, // random small distance
		},
		{
			coordinate1: Coordinates{
				Latitude:  53.339428,
				Longitude: -6.257664,
			},
			coordinate2: Coordinates{
				Latitude:  53.4692815,
				Longitude: -9.436036,
			},
			greatCircleDistance: 211.1720522962854, // random large distance
		},
	}

	greatCircleCalculator := NewGreatCircleCalculator()

	for _, testCase := range testCases {
		greatCircleDistance := greatCircleCalculator.Calculate(testCase.coordinate1, testCase.coordinate2)
		assert.Equal(t, testCase.greatCircleDistance, greatCircleDistance)
	}
}
