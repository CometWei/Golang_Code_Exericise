package main

import "fmt"

func main() {
	var in, ten, five, one int
	fmt.Scanf("%d", &in)
	ten = in / 10
	in = in % 10
	five = in / 5
	in = in % 5
	one = in / 1
	fmt.Printf("%d\n", ten)
	fmt.Printf("%d\n", five)
	fmt.Printf("%d\n", one)
}
