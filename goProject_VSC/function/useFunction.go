package main

import "fmt"

func average(a []float64) float64 {
	var total float64 = 0
	for _, value := range a {
		total += value
	}
	return total / float64(len(a))
}
func runAverage() float64 {
	return average(mygrade)
}

var mygrade = []float64{98, 88, 75, 68, 72}

func main() {

	fmt.Println("средняя оценка - ", runAverage())
}
