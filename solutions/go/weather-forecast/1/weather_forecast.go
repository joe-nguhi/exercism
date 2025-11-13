// Package weather provides weather forecasting functionality.
package weather

var (
    // CurrentCondition stores the current weather condition.
	CurrentCondition string
    // CurrentLocation Stores the name of the city for which the weather is reported.
	CurrentLocation  string
)

// Forecast sets the package-level variables CurrentLocation and CurrentCondition to the provided values and returns a formatted string describing the weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
