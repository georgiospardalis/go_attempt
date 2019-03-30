package ride

import (
	"math/big"
	"testing"
)

func TestCheckPositionsInvalid(t *testing.T) {
	distanceTraveled := 200.00
	timeDelta := 1.0

	if checkPositions(distanceTraveled, timeDelta) {
		t.Error("Should not return ok for such a great distance in such a short time")
	}
}

func TestCheckPositionsValid(t *testing.T) {
	distanceTraveled := 10.00
	timeDelta := 1.0

	if !checkPositions(distanceTraveled, timeDelta) {
		t.Error("Should return ok")
	}
}

func TestDistanceInKm(t *testing.T) {
	expected := big.NewFloat(0.0005536352197098919)

	positionA := NewPosition(1.00, 1.00, 1553984501)
	positionB := NewPosition(1.00, 1.000005, 1553984502)

	distance := distanceInKm(positionB, positionA)
	distanceCalculated := big.NewFloat(distance)

	if expected.Cmp(distanceCalculated) != 0 {
		t.Errorf("Expected distance: %f, got: %f", 0.0005536352197098919, distance)
	}
}

func TestWasHighRate(t *testing.T) {
	tsHigh := int64(1483236000)

	if !wasHighRate(tsHigh){
		t.Error("Timestamp in High Rate was not marked accordingly...")
	}

	tsLow := int64(1483268400)

	if wasHighRate(tsLow) {
		t.Error("Timestamp in Low Rate was not marked accordingly...")
	}
}

func TestWasStationary(t *testing.T) {
	timeDeltaInHours := 1.00
	distanceCovered := 0.00

	if !wasStationary(distanceCovered, timeDeltaInHours) {
		t.Error("Stationary time was not marked accordingly...")
	}
}

func TestTimeDeltaInHours(t *testing.T) {
	tsA := int64(1483268400)
	tsB := int64(1483275600)

	if timeDeltaInHours(tsB, tsA) != 2.00 {
		t.Error("Time difference should be 2.00 hours...")
	}
}