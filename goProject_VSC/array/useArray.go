package main

import "fmt"

func main() {
	x := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}
	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
	fmt.Println("--создание срезов--")
	arr := []int{1, 2, 4, 5, 6, 7, 8, 9} //создание массива
	a := arr[0:3]                        //создание среза из массива arr
	fmt.Println("массив arr ", arr)
	fmt.Println("slice(срез[0:3]) a на основе массива arr", a)
	fmt.Println("-----------------------------")
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	slice3 := slice1[:2]
	copy(slice2, slice1)
	fmt.Println(slice1, slice2, slice3)
}
