package jedlik
import (
    "math"
    "fmt"
)

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
    turnNeeded := int(math.Ceil(float64(trackDistance)/float64(car.speed)))
    return car.battery >= turnNeeded*car.batteryDrain
}
