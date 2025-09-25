package speed

// Car is a remote controlled car with a certain speed, battery, battery drain per drive and
// a distance driven so far.
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
		distance:     0,
	}
}

// Track is a race track with a certain distance.
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	// Check if the car has enough battery to drive
	if car.battery < car.batteryDrain {
		return car
	}

	car.battery -= car.batteryDrain
	car.distance += car.speed

	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	possibleDistance := car.speed * (car.battery / car.batteryDrain)
	return possibleDistance >= track.distance
}
