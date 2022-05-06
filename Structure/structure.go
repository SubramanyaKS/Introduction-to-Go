package main

import "fmt"

type Address struct {
	Name    string
	city    string
	Pincode int
}

func main() {
	var a Address
	fmt.Println(a)

	a1 := Address{"Subramanya", "Shimoga", 577201}

	fmt.Println("Address1: ", a1)
	a2 := Address{Name: "Delhi"}
	fmt.Println("Address2: ", a2)
}
