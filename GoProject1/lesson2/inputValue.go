package main

import "fmt"

func main() {
	fmt.Println("enter a temperature in farengate")
	var far float64
	fmt.Scanf("%f", &far)
	output := (far - 32) * 5 / 9
	var c = fmt.Sprintf("temperature in celsium = %f", output)
	fmt.Println(c)
	var user = []string{"Ivan", "Anton", "Natalia", "Maria"}
	for _, value := range user {
		fmt.Println(value)
	}
	{
		for i := 1; i < 10; i++ {
			for k := 1; k < 10; k++ {
				fmt.Print(i*k, "\t")
			}
			fmt.Println()
		}
	}
}
