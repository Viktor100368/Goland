package main

import "fmt"

func someArgument(arg ...int) int {
	total := 0
	for _, value := range arg {
		total += value
	}
	return total
}

func main() {
	a := someArgument(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("total = ", a)
}
