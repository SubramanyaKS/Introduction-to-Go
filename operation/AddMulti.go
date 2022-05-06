package main

import "fmt"

func main() {
	a := addTwoNumbers(3, 5)
	fmt.Println(a)
	m := multiplyTwoNumbers(5, 9)
	fmt.Println(m)

}
func addTwoNumbers(x, y int) int {
	sum := x + y
	return sum
}
func multiplyTwoNumbers(x, y int) int {
	product := x * y
	return product
}
