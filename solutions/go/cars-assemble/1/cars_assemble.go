package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var carsPerHour float64

    carsPerHour = float64(productionRate) * successRate/100
    return carsPerHour
    
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var carsPerMinute float64

    carsPerMinute = (float64(productionRate)/60) * (successRate/100)
    return int(carsPerMinute)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var batchCostPerUnit float64
    var singleCostPerUnit float64
    var singles uint
    var batches uint

    batchCostPerUnit = 9500.00
    singleCostPerUnit = 10000.00

    batches = (uint(carsCount) / 10)
    singles = (uint(carsCount) % 10)
    if carsCount < 10 {
        return singles * uint(singleCostPerUnit)
    }
    if batches == 0 {
        return uint(carsCount) * uint(batchCostPerUnit)
    }
    
    return singles * uint(singleCostPerUnit) + batches * 10 * uint(batchCostPerUnit)
    
}
