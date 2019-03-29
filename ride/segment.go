package ride

type Segment struct {
	highRate   bool
	stationary bool
	distance  float64
}

func segmentForPositions(positionB, positionA Position) (bool, Segment) {
	var segment Segment

	distanceTraveled := distanceInKm(positionB, positionA)
	timeDelta := timeDeltaInHours(positionB.ts, positionA.ts)

	if !checkPositions(distanceTraveled, timeDelta) {
		return false, segment
	}

	segment = Segment{
		wasHighRate(positionA.ts),
		wasStationary(distanceTraveled, timeDelta),
		distanceTraveled}

	return true, segment
}