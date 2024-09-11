package main

import "fmt"

func main() {
	type NoKTP string
	type Maried bool

	var noKTPdimas NoKTP = "039890398493098"
	var status Maried = false

	fmt.Println(noKTPdimas)
	fmt.Println(status)
}
