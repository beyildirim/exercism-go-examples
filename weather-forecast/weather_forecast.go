// Package weather loads weather package.
package weather

// CurrentCondition keeps the temperature.
var CurrentCondition string
// CurrentLocation keeps the location.
var CurrentLocation string

// Forecast returns the weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
