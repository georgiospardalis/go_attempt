package haversine

import (
	"math/big"
	"testing"
)

func TestHaversineDistance(t *testing.T) {
	expectedDistance := big.NewFloat(0.005389130468647274)
	haversineResult := HaversineDistance(
		37.966627,
		23.728263,
		37.966660,
		23.728308)
	distance := big.NewFloat(haversineResult)

	if expectedDistance.Cmp(distance) != 0 {
		t.Errorf("Wrong estimation. Expected: %f, got: %f", 0.005389130468647274, haversineResult)
	}
}
