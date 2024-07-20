package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	prodRateF := float64(productionRate)
	var carsPerHour float64 = prodRateF * successRate
	return carsPerHour / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var carsPerHour = CalculateWorkingCarsPerHour(productionRate, successRate)
	carsPerHourI := int(carsPerHour)
	return carsPerHourI / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	carsCountU := uint(carsCount)
	var tens = carsCountU / 10
	var ones = carsCountU % 10
	return (tens * 95000) + (ones * 10000)
}
