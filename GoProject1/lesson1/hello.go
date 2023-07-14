package main

import "fmt"

func main() {
	fmt.Println("hello dolly")
	var age = 55
	var name = "Viktor"
	var c = fmt.Sprintf("my name is %s, and my age is %d", name, age)
	fmt.Println(c)
}
