package jedlik
import "fmt"

func (car *Car) Drive() {
    if (car.batteryDrain <= car.battery){
        car.distance += car.speed
        car.battery -= car.batteryDrain
    }
}

func (car Car) DisplayDistance() string {
    return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car Car) DisplayBattery() string {
    return fmt.Sprintf("Battery at %d%%", car.battery)
}

func (car Car) CanFinish(trackDistance int) bool {
    turnNeeded := trackDistance/car.speed
    if (trackDistance % car.speed != 0){
        turnNeeded++
    }
    return car.battery >= turnNeeded*car.batteryDrain
}
