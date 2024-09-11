package main

import "fmt"

func main() {
	var name string

	name = "Dimas Maulana" // pemakaian terdefinisi tipe data ke 1
	fmt.Println(name)

	name = "Dimas" // pemakaian terdefinisi tipe data  ke 2
	fmt.Println(name)

	var friendName = "Budi" // pemakaian tidak menggunakan definisi tipe data ke 3
	fmt.Println(friendName)

	var age = 30 // pemakaian tidak menggunakan definisi tipe data ke 4
	fmt.Println(age)

	firstName := "Dimdevs" // Represenstasi Langsung awal dan wajib
	fmt.Println(firstName)

	firstName = "Dimas"
	fmt.Println(firstName)

	var (
		kesatu = "Dimas"
		kedua  = "Maulana"
	)

	fmt.Println(kesatu)
	fmt.Println(kedua)
}
