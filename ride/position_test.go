package ride

import "testing"

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