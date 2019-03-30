package ride

import (
	"github.com/georgiospardalis/beat_assignment/haversine"
	"time"
)

func checkPositions(distanceTraveled, timeDelta float64) bool {
	if distanceTraveled / timeDelta >= 100 {
		return false
	}

	return true
}

func distanceInKm(positionB, positionA Position) float64 {
	return haversine.HaversineDistance(positionB.latitude, positionB.longitude, positionA.latitude, positionA.longitude)
}

func timeDeltaInHours(timestampB, timestampA int64) float64 {
	timeA := time.Unix(timestampA,0)
	timeB := time.Unix(timestampB, 0)

	diff := timeB.Sub(timeA)

	return diff.Hours()
}

func wasHighRate(startTimestamp int64) bool {
	startTime := time.Unix(startTimestamp, 0)

	if startTime.Hour() <= 5 {
		return true
	}

	return false
}

func wasStationary(distanceTraveled, timeDelta float64) bool {
	if distanceTraveled == 0 && timeDelta >= 1.00 {
		return true
	}

	return false
}