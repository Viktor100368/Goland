package main

import "fmt"

func main() {
	x := 0
	inc := func() int {
		x += 1
		return x
	}
	dec := func() int {
		x -= 1
		return x
	}
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(dec())
	fmt.Println(dec())
	fmt.Println(dec())
	fmt.Println(dec())
}
