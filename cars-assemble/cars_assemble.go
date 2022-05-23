package cars

const costOnly1Car = 10000
const cost10Cars = 95000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	if carsCount < 10 {
		return uint(carsCount * costOnly1Car)
	} else {
		carModedBy10 := carsCount % 10
		carDividedBy10 := carsCount / 10
		return uint(carDividedBy10*cost10Cars) + uint(carModedBy10*costOnly1Car)
	}
}
