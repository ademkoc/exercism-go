// Package weather provides a forecast info.
package weather

var (
    // CurrentCondition provides weather condition.
	CurrentCondition string
    // CurrentLocation provides city location.
	CurrentLocation  string
)

// Forecast provides weather condition for provided city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
