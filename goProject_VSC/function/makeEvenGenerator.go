package main

import "fmt"

func evenGenerator() func() uint {
	//создаем замыкание
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
func main() {
	// для использования замыкания, нужно присвоить значение функции переменной
	//вызвать переменную как функцию
	nextEven := evenGenerator()
	for i := 0; i < 10; i++ {
		fmt.Println(nextEven())
	}
}
