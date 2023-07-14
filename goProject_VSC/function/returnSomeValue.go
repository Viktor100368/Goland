package main

import "fmt"

func someValue() (string, string) {
	return "Hello", "Dolly"
}

func main() {
	a, b := someValue()
	fmt.Println(a, b)
}
