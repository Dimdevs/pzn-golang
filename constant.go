package main

import "fmt"

func main() {
	const firstName string = "DIMAS"
	const lastname = "DIMAS"
	const value = 100

	fmt.Println(firstName)
	fmt.Println(lastname)
	fmt.Println(value)

	// Multiple Constant
	const (
		kesatu string = "DIMAS"
		kedua         = "Maulana"
		ketiga        = 100
	)

	fmt.Println(kesatu)
	fmt.Println(kedua)
	fmt.Println(ketiga)

}
