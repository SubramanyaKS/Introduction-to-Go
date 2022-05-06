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
	//Acessing member
	fmt.Println("Name: ", a1.Name)

	a2 := Address{city: "Delhi"}
	fmt.Println("Address2: ", a2)
	a2.Name = "Pratheek"
	fmt.Println("Adress2 (modifier): ", a2)
}
