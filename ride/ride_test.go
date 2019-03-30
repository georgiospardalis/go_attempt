package ride

import (
	"testing"
)

func TestNewPosition(t *testing.T) {
	position := NewPosition(1.00, 1.00, 1553984501)

	if position.ts != 1553984501 {
		t.Errorf("Expected ts: %d, got: %d", 1553984501, position.ts)
	}

	if position.longitude != 1.00 {
		t.Errorf("Expected longituted: %f, got: %f", 1.00, position.longitude)
	}

	if position.latitude != 1.00 {
		t.Errorf("Expected latitude: %f, got: %f", 1.00, position.latitude)
	}
}

func TestNewRide(t *testing.T) {
	var positions []Position

	position := NewPosition(1.00, 1.00, 1553984501)
	positions = append(positions, position)

	cabRide := NewRide(1, positions)

	if cabRide.id != 1 {
		t.Errorf("Expected id: %d, got: %d", 1, cabRide.id)
	}

	sliceLen := len(cabRide.positions)

	if sliceLen != 1 {
		t.Errorf("Expected Positions slice to have a lenght of: %d, got: %d", 1, sliceLen)
	}
}

func TestRide_Id(t *testing.T) {
	cabRide := NewRide(1, nil)
	gotId := cabRide.Id()

	if gotId != 1 {
		t.Errorf("Expected Id() to return value of: %d, got: %d", 1, gotId)
	}
}

func TestRide_PrepareSegments(t *testing.T) {
	var positions []Position

	position := NewPosition(1.00, 1.00, 1553984501)
	positions = append(positions, position)
	position = NewPosition(30.00, 30.00, 1553984505)
	positions = append(positions, position)
	position = NewPosition(1.00, 1.0000005, 1553984502)
	positions = append(positions, position)

	cabRide := NewRide(1, positions)
	cabRide.PrepareSegments()

	numOfSegments := len(cabRide.segments)

	if numOfSegments > 1 {
		t.Errorf("Expected 1 segment after cleanup, got %d", numOfSegments)
	}
}

func TestRide_EstimateFare(t *testing.T) {
	var positions []Position

	position := NewPosition(1.00, 1.00, 1553984501)
	positions = append(positions, position)
	position = NewPosition(1.00, 1.00, 1553984505)
	positions = append(positions, position)

	cabRide := NewRide(1, positions)
	cabRide.PrepareSegments()

	calculatedFare := cabRide.EstimateFare()

	if calculatedFare != 3.47 {
		t.Errorf("Expected fare %f: got: %f", 3.47, calculatedFare)
	}
}