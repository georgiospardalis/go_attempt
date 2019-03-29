package ride

const startingFare = 1.30
const minimumFare = 3.47
const normalRate = 0.74
const highRate = 1.30
const stationaryRate = 11.90

type ride interface {
	PrepareSegments()
	EstimateFare() float64
}

type Ride struct {
	id        int
	positions []Position
	segments  []Segment
}

func NewRide(id int, ridePositions []Position) Ride {
	return Ride{id: id, positions: ridePositions}
}

func (ride Ride) Id() int {
	return ride.id
}

func (ride *Ride) PrepareSegments() {
	var positionA = ride.positions[0]
	var rideSegments []Segment

	for index, positionB := range ride.positions[1:] {
		ok, segment := segmentForPositions(positionB, positionA)

		if ok {
			rideSegments = append(ride.segments, segment)
			positionA = ride.positions[index]
		}
	}

	ride.segments = rideSegments
}

func (ride Ride) EstimateFare() float64 {
	sum := startingFare

	for _, segment := range ride.segments {
		if segment.stationary {
			sum += stationaryRate
		} else if segment.highRate {
			sum += segment.distance * highRate
		} else {
			sum += segment.distance * normalRate
		}
	}

	if sum < minimumFare {
		return minimumFare
	}

	return sum
}