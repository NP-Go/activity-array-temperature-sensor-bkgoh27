package main

import "fmt"

func main() {
	temperature := [24]float64{20.1, 24, 27.3, 30.1, 26.4, 22.2, 20.1, 24, 27.3, 30.1, 26.4, 20.1, 24, 27.3, 30.1, 26.4, 20.1, 24, 27.3, 30.1, 26.4, 20.1, 24, 27.3}

	var sum float64
	for i := 0; i < len(temperature); i++ {
		sum += temperature[i]
	}
	averageTemp := sum / float64(len(temperature))
	fmt.Println("average temperature", averageTemp)
}
