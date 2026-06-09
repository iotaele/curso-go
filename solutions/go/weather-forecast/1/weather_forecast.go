// Package weather provides information about the current condition and the weather forecast.
package weather

var (
    // CurrentCondition stores weather current condition.
	CurrentCondition string
    // CurrentLocation stores weather current location.
	CurrentLocation  string
)

// Forecast returns a string with a message describing the weather condition at the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
