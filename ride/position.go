package ride

type Position struct {
	latitude float64
	longitude float64
	ts int64
}

func NewPosition(latitude, longitude float64, ts int64) Position {
	return Position{latitude: latitude, longitude: longitude, ts: ts}
}