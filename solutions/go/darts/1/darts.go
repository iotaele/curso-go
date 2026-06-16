package darts

import "math"

func Score(x, y float64) int {
    absX := math.Abs(x)
    absY := math.Abs(y)
    v := math.Sqrt(math.Pow(absX, 2.0) + math.Pow(absY, 2.0))
	if v <= 1 {
        return 10
    } else if v <= 5 {
        return 5
    } else if v <= 10 {
        return 1
    }
    return 0
}
