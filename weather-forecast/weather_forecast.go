// Package weather have function Forecast() to tell current condition on specific city.
package weather

// CurrentCondition define current weather condition.
var CurrentCondition string

// CurrentLocation define city or location name.
var CurrentLocation string

// Forecast() method will return weather condition and location based on parameter city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
