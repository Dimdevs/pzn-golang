package main

import "fmt"

func main() {
	var name1 = "Dimas"
	var name2 = "Maulana"
	var result bool = name1 == name2

	fmt.Println(result)

	var value1 = 101
	var value2 = 100
	var result2 bool = value1 > value2

	fmt.Println(result2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
