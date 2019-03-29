package haversine

import (
	"math"
)

const rEarth = 6372.8 // in km

type radianPosition struct {
	latitude  float64 // latitude, radians
	longitude float64 // longitude, radians
}

func radianPositionFrom(latitude, longitude float64) radianPosition {
	return radianPosition{latitude * math.Pi / 180, longitude * math.Pi / 180}
}

func haversine(theta float64) float64 {
	return .5 * (1 - math.Cos(theta))
}

func haversineFormula(p1, p2 radianPosition) float64 {
	return 2 * rEarth * math.Asin(math.Sqrt(haversine(p2.latitude-p1.latitude)+
		math.Cos(p1.latitude)*math.Cos(p2.latitude)*haversine(p2.longitude-p1.longitude)))
}

func HaversineDistance(finalLatitude, finalLongitude, startLatitude, startLongitude float64) float64 {
	finalPosition := radianPositionFrom(finalLatitude, finalLongitude)
	startingPosition := radianPositionFrom(startLatitude, startLongitude)

	return haversineFormula(finalPosition, startingPosition)
}