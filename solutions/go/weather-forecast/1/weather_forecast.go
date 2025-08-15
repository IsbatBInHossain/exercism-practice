// Package weather forecast the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition represents the current condition.
var CurrentCondition string
// CurrentLocation represents the current city.
var CurrentLocation string 

// Forecast returns a string describing the current weather condition in the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
