// Package weather provides tools to forecast weather.
package weather

// CurrentCondition represents current weather conditions.
var CurrentCondition string

// CurrentLocation represents current location.
var CurrentLocation string

// Forecast returns a string value of weather condition for a location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
