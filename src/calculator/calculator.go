package calculator

import "math"

const earthRadiusInKm = 6371

// Calculator defines the interface for calculating a value from two floating-point numbers
type Calculator interface {
	Calculate(Coordinates, Coordinates) float64
}

// Coordinates defines the structure of a geographical location
type Coordinates struct {
	Latitude  float64
	Longitude float64
}

type greatCircleCalculator struct{}

// NewGreatCircleCalculator creates and returns a calculator for great-circle distance
func NewGreatCircleCalculator() Calculator {
	return &greatCircleCalculator{}
}

// Calculate calculates the great-circle distance between two coordinates
//
// Ref: https://en.wikipedia.org/wiki/Great-circle_distance
// Based on: "A more complicated formula that is accurate for all distances is the following special case of the Vincenty formula for an ellipsoid with equal major and minor axes"
func (c *greatCircleCalculator) Calculate(coordinate1, coordinate2 Coordinates) float64 {
	coordinate1Latitude := degreesToRadians(coordinate1.Latitude)
	coordinate1LatitudeSine := math.Sin(coordinate1Latitude)
	coordinate1LatitudeCosine := math.Cos(coordinate1Latitude)

	coordinate1Longitude := degreesToRadians(coordinate1.Longitude)

	coordinate2Latitude := degreesToRadians(coordinate2.Latitude)
	coordinate2LatitudeSine := math.Sin(coordinate2Latitude)
	coordinate2LatitudeCosine := math.Cos(coordinate2Latitude)

	coordinate2Longitude := degreesToRadians(coordinate2.Longitude)

	deltaLongitude := coordinate1Longitude - coordinate2Longitude
	deltaLongitudeSine := math.Sin(deltaLongitude)
	deltaLongitudeCosine := math.Cos(deltaLongitude)

	yPart1 := coordinate2LatitudeCosine * deltaLongitudeSine
	yPart2 := (coordinate1LatitudeCosine * coordinate2LatitudeSine) - (coordinate1LatitudeSine * coordinate2LatitudeCosine * deltaLongitudeCosine)
	y := math.Sqrt(math.Pow(yPart1, 2) + math.Pow(yPart2, 2))

	x := (coordinate1LatitudeSine * coordinate2LatitudeSine) + (coordinate1LatitudeCosine * coordinate2LatitudeCosine * deltaLongitudeCosine)

	return math.Atan2(y, x) * earthRadiusInKm
}

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}
