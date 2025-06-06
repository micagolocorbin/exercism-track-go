package speed

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
		speed:        speed,
		batteryDrain: batteryDrain,
	}
}

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
	if car.battery-car.batteryDrain < 0 {
		return car
	} else {
		return Car{
			distance:     car.distance + car.speed,
			battery:      car.battery - car.batteryDrain,
			speed:        car.speed,
			batteryDrain: car.batteryDrain,
		}
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	speedPerDist := track.distance / car.speed
	endurance := speedPerDist * car.batteryDrain
	return endurance <= car.battery
}
