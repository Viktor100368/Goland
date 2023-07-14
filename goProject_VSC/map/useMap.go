package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["one"] = 1
	x["two"] = 2
	x["three"] = 3
	x["four"] = 4
	x["five"] = 5
	for key, _ := range x {
		fmt.Println(x[key])
	}
	//другой способ создания карты
	fmt.Println("----------------------------")
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Litium",
		"Be": "Berilium",
	}
	for key, _ := range elements {
		fmt.Println(key, elements[key])
	}
	fmt.Println("-----------------------------")
	//создание карты карт
	elementDescription := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Litium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Berilium",
			"state": "solid",
		},
	}
	for key, _ := range elementDescription {
		fmt.Println(key, elementDescription[key])
	}
}
