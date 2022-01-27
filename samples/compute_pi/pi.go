package main

import "math/rand"

// computePi computes pi decimals using Monte Carlo simulation
func computePi(n int) float64 {
	var total int = 0
	var totalIn int = 0
	for i := 0; i < n; i++ {
		var x = rand.Float64()
		var y = rand.Float64()
		if x*x+y*y < 1 {
			totalIn++
		}
		total++
	}
	return float64(totalIn) * 4.0 / float64(total)
}

func main() {
	computePi(5000)
}
