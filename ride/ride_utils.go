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
	timestampDiff := timestampB - timestampA
	diffTime := time.Unix(timestampDiff, 0)

	return float64(diffTime.Second()) / 3600.00
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