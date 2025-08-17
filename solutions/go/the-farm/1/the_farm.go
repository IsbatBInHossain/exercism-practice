package thefarm
import (
    "fmt"
    "errors"
)

// TODO: define the 'DivideFood' function
func DivideFood (fc FodderCalculator, cows int) (float64, error) {
    fodderAmount, err := fc.FodderAmount(cows)
    if err != nil {
        return 0.0 ,err
    }
    fatteningFactor, err := fc.FatteningFactor()
    if err != nil {
        return 0.0, err
    }
    return fodderAmount * fatteningFactor / float64(cows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood (fc FodderCalculator, cows int) (float64, error) {
    if cows <= 0 {
        return 0.0, errors.New("invalid number of cows")
    }
    return DivideFood(fc, cows)
}

type InvalidCowsError struct {
    cows int
    message string
}

func (e *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}
// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error {
    switch {
    case cows < 0:
        return &InvalidCowsError{cows: cows, message: "there are no negative cows"}
    case cows == 0:
        return &InvalidCowsError{cows: cows, message: "no cows don't need food"}
    default:
    	return nil
    }
}

