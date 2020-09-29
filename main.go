package main

import "fmt"

func main() {
	var base int
	var altura int

	fmt.Scanln(&base)
	fmt.Scanln(&altura)
	area := (base * altura) / 2

	fmt.Println(area)
}
