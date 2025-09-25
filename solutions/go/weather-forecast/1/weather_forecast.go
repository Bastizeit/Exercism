// Package weather provides weather forecast functionalities.
package weather

// CurrentCondition holds the current weather condition.
var CurrentCondition string
// CurrentLocation holds the current location for the weather forecast.
var CurrentLocation string

// Forecast returns a formatted string with the current location and weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
