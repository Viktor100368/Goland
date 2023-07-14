package main

import "fmt"

var array = []int{-1, -43, 5, -24, 4, 6, 0, -43}

func main() {
	for _, value := range array {
		if value < 0 {
			continue
		} else {
			switch value {
			case 0:
				fmt.Println("zero")
			case 4:
				fmt.Println("four")
			case 5:
				fmt.Println("five")
			case 6:
				fmt.Println("six")
			default:
				fmt.Println("no number")
			}
		}
	}
}
