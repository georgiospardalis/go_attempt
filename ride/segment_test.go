package ride

import (
	"testing"
)

func TestSegmentForPositionsValid(t *testing.T) {
	positionA := NewPosition(1.00, 1.00, 1553984501)
	positionB := NewPosition(1.00, 1.00, 1553984535)

	successful, _ := segmentForPositions(positionB, positionA)

	if !successful {
		t.Error("Should get a valid segment")
	}
}

func TestSegmentForPositionsInvalid(t *testing.T) {
	positionA := NewPosition(1.00, 1.00, 1553984501)
	positionB := NewPosition(30.00, 30.00, 1553984535)

	successful, _ := segmentForPositions(positionB, positionA)

	if successful {
		t.Error("Should not get a valid segment")
	}
}